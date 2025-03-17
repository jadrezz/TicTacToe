package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

type cell struct {
	id    int
	value rune
}

func drawBoard(board []cell) {
	fmt.Println(strings.Repeat("_", 11))
	for n, v := range board {
		fmt.Printf("|%c| ", v.value)
		if (n+1)%3 == 0 {
			fmt.Println()
			continue
		}
	}
	fmt.Println(strings.Repeat("_", 11))
}

func makeStep(board []cell, num int, playerFigure rune) bool {
	if num < 0 || num > 9 || board[num].value == 'X' || board[num].value == 'Y' {
		return false
	}
	board[num].value = playerFigure
	return true
}

func isWin(board []cell) bool {
	for _, combo := range winCombinations {
		posA, posB, posC := combo[0], combo[1], combo[2]
		if board[posA].value == board[posB].value && board[posB].value == board[posC].value {
			return true
		}
	}
	return false
}

func isDraw(board []cell) bool {
	for _, set := range board {
		if unicode.IsNumber(set.value) {
			return false
		}
	}
	return true
}

func main() {
	board := []cell{
		{1, '1'}, {3, '2'}, {3, '3'},
		{4, '4'}, {5, '5'}, {6, '6'},
		{7, '7'}, {8, '8'}, {9, '9'},
	}
	player := 'X'
	input := bufio.NewScanner(os.Stdin)

	for {
		drawBoard(board)
		fmt.Printf("Игрок %c, сделайте ход\n", player)

		input.Scan()
		n, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println("Некорректный ввод")
			continue
		}

		if !makeStep(board, n-1, player) {
			fmt.Println("Выход за пределы доски или занятая клетка")
			continue
		}

		if isWin(board) {
			drawBoard(board)
			fmt.Printf("Победили %c\n", player)
			break
		}

		if isDraw(board) {
			drawBoard(board)
			fmt.Println("Ничья!")
			break
		}

		player = map[rune]rune{'X': 'Y', 'Y': 'X'}[player]

	}
}
