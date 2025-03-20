package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	red     string = "\033[0;31m%s\033[0m"
	yellow  string = "\033[1;33m%s\033[0m"
	green   string = "\033[0;32m%s\033[0m"
	cyan    string = "\033[0;36m%s\033[0m"
	xplayer string = "\033[0;94m%s\033[0m"
	oplayer string = "\033[0;95m%s\033[0m"
)

var winCombinations = [][]int{
	{0, 1, 2}, // Первая строка
	{3, 4, 5}, // Вторая строка
	{6, 7, 8}, // Третья строка
	{0, 3, 6}, // Первый столбец
	{1, 4, 7}, // Второй столбец
	{2, 5, 8}, // Третий столбец
	{0, 4, 8}, // Главная диагональ
	{2, 4, 6}, // Побочная диагональ
}

func drawBoard(board []rune) {
	fmt.Println(strings.Repeat("_", 11))
	for n, v := range board {
		switch v {
		case 'X':
			fmt.Printf(toColor("xplayer", fmt.Sprintf("|%c| ", v)))
		case 'O':
			fmt.Printf(toColor("oplayer", fmt.Sprintf("|%c| ", v)))
		default:
			fmt.Printf("|%c| ", v)
		}
		if (n+1)%3 == 0 {
			fmt.Println()
			continue
		}
	}
	fmt.Println(strings.Repeat("_", 11))
}

func toColor(color string, msg string) string {
	switch color {
	case "cyan":
		return fmt.Sprintf(cyan, msg)
	case "red":
		return fmt.Sprintf(red, msg)
	case "green":
		return fmt.Sprintf(green, msg)
	case "yellow":
		return fmt.Sprintf(yellow, msg)
	case "xplayer":
		return fmt.Sprintf(xplayer, msg)
	case "oplayer":
		return fmt.Sprintf(oplayer, msg)
	default:
		return msg
	}
}

func makeStep(board []rune, num int, playerFigure rune) bool {
	if (num < 0 || num >= 9) || (board[num] == 'X' || board[num] == 'O') {
		return false
	}
	board[num] = playerFigure
	return true
}

func isWin(board []rune) bool {
	for _, combo := range winCombinations {
		posA, posB, posC := combo[0], combo[1], combo[2]
		if board[posA] == board[posB] && board[posB] == board[posC] {
			return true
		}
	}
	return false
}

func isDraw(board []rune) bool {
	for _, val := range board {
		if unicode.IsDigit(val) {
			return false
		}
	}
	return true
}

func main() {
	board := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	player := 'X'
	input := bufio.NewScanner(os.Stdin)

	for {
		drawBoard(board)
		fmt.Println(toColor("cyan", fmt.Sprintf("Игрок %c, сделайте ход", player)))

		input.Scan()
		n, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println(toColor("red", "Некорректный ввод"))
			continue
		}

		if !makeStep(board, n-1, player) {
			fmt.Println(toColor("red", "Выход за пределы доски или занятая клетка"))
			continue
		}

		if isWin(board) {
			drawBoard(board)
			fmt.Println(toColor("green", fmt.Sprintf("Победили %c\n", player)))
			break
		}

		if isDraw(board) {
			drawBoard(board)
			fmt.Println(toColor("yellow", "Ничья!"))
			break
		}

		player = map[rune]rune{'X': 'O', 'O': 'X'}[player]

	}
}
