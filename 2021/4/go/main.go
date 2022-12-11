package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	board   [][]string
	visited [][]bool
}

func NewBoard(board [][]string) *Board {
	temp := Board{}
	temp.board = board
	temp.visited = make([][]bool, 5)
	for i := 0; i < 5; i++ {
		temp.visited[i] = make([]bool, 5)
	}

	return &temp
}

func (b *Board) Draw(draw string) {
	for i, row := range b.board {
		for j, cell := range row {
			if draw == cell {
				b.visited[i][j] = true
			}
		}
	}
}

func (b *Board) SumUnvisited() int {
	val := 0

	for i, row := range b.visited {
		for j, visited := range row {
			if !visited {
				temp, _ := strconv.ParseInt(b.board[i][j], 10, 64)
				val += int(temp)
			}
		}
	}

	return val
}

func (b *Board) HasWon() bool {
	win := false

	for i := 0; i < 5; i++ {
		// Do we have any wins vertically?
		if b.visited[0][i] && b.visited[1][i] && b.visited[2][i] && b.visited[3][i] && b.visited[4][i] {
			win = true
		}

		// Do we have any wins horizontally?
		if b.visited[i][0] && b.visited[i][1] && b.visited[i][2] && b.visited[i][3] && b.visited[i][4] {
			win = true
		}
	}

	return win
}

func (b *Board) String() string {
	var board string

	for _, row := range b.board {
		for _, cell := range row {
			board += fmt.Sprintf("%s\t", cell)
		}
		board += fmt.Sprintf("\n")
	}

	return board
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Parse Turns
	scanner.Scan()
	turns := strings.Split(scanner.Text(), ",")

	// Empty Line
	scanner.Scan()

	// Parse boards
	boards := make([]*Board, 0)
	boards2 := make([]*Board, 0)

	for scanner.Scan() {
		board := make([][]string, 5)
		for i := 0; i < 5; i++ {
			board[i] = make([]string, 5)
		}

		// Parse Board Row 1
		fmt.Sscan(scanner.Text(), &board[0][0], &board[0][1], &board[0][2], &board[0][3], &board[0][4])

		// Parse Board Row 2
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &board[1][0], &board[1][1], &board[1][2], &board[1][3], &board[1][4])

		// Parse Board Row 3
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &board[2][0], &board[2][1], &board[2][2], &board[2][3], &board[2][4])

		// Parse Board Row 4
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &board[3][0], &board[3][1], &board[3][2], &board[3][3], &board[3][4])

		// Parse Board Row 5
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &board[4][0], &board[4][1], &board[4][2], &board[4][3], &board[4][4])

		// Empty Line
		scanner.Scan()

		boards = append(boards, NewBoard(board))
		boards2 = append(boards2, NewBoard(board))
	}

	part1 := 0

sim:
	for _, turn := range turns {
		for i := 0; i < len(boards); i++ {
			boards[i].Draw(turn)

			if boards[i].HasWon() {
				val, _ := strconv.ParseInt(turn, 10, 64)
				part1 = int(val) * boards[i].SumUnvisited()
				break sim
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)

	winningBoards := make(map[int]bool, 0)
	part2 := 0

sim2:
	for _, turn := range turns {
		for i := 0; i < len(boards2); i++ {
			boards2[i].Draw(turn)

			if boards2[i].HasWon() {
				winningBoards[i] = true

				if len(winningBoards) == len(boards2) {
					val, _ := strconv.ParseInt(turn, 10, 64)
					part2 = int(val) * boards2[i].SumUnvisited()
					break sim2
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", part2)
}
