package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

const boardCells int = 9

type cell struct {
	cellNum int
	player  string
}

func showArea(board *[boardCells]cell) {
	for n, cell := range board {
		fmt.Printf("%d: %s\t", cell.cellNum, cell.player)
		if (n+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func makeStep(board *[boardCells]cell, num int, figure string) bool {
	if num < 0 || num >= boardCells || board[num].player != "" {
		fmt.Println("Клетка занята или вы вышли за поле")
		return false
	}
	board[num].player = figure
	return true
}

func checkWin(board *[boardCells]cell) (win bool, player string) {
	for _, combo := range winCombinations {
		a, b, c := combo[0], combo[1], combo[2]
		if board[a].player != "" &&
			board[a].player == board[b].player &&
			board[b].player == board[c].player {
			return true, board[a].player
		}
	}
	return false, ""
}

func isDraw(board *[boardCells]cell) bool {
	for _, cell := range board {
		if cell.player == "" {
			return false
		}
	}
	return true
}

func main() {
	board := [boardCells]cell{
		{cellNum: 1}, {cellNum: 2}, {cellNum: 3},
		{cellNum: 4}, {cellNum: 5}, {cellNum: 6},
		{cellNum: 7}, {cellNum: 8}, {cellNum: 9},
	}
	gamer := "X"
	input := bufio.NewScanner(os.Stdin)

	for {
		showArea(&board)
		fmt.Printf("Игрок %s, выберите цифру для хода\n", gamer)
		input.Scan()
		num, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println("Некорректный ввод. Повторите попытку")
			continue
		}
		if !makeStep(&board, num-1, gamer) {
			continue
		}

		if win, player := checkWin(&board); win {
			fmt.Printf("Победил игрок %s\n", player)
			fmt.Println("Игра окончена")
			break
		} else if isDraw(&board) {
			fmt.Println("Ничья! Игра окончена.")
			break
		}

		gamer = map[string]string{"X": "Y", "Y": "X"}[gamer]
	}

}
