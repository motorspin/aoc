package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	inputs := make([]string, 0)

	for scanner.Scan() {
		row := scanner.Text()
		inputs = append(inputs, row)
	}

	grid := make([][]int, len(inputs))
	visible := make([][]bool, len(inputs))

	for index, row := range inputs {
		for _, char := range row {
			val, _ := strconv.ParseInt(string(char), 10, 64)
			grid[index] = append(grid[index], int(val))
			visible[index] = append(visible[index], false)
		}
	}

	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs); j++ {
			// This initializes the edges
			if i == 0 || j == 0 || i == len(inputs)-1 || j == len(inputs)-1 {
				visible[i][j] = true
				continue
			}
		}
	}

	// Compute going downward
	for i := 1; i < len(inputs)-1; i++ {
		height := grid[i][0]
		for j := 1; j < len(inputs)-1; j++ {
			//fmt.Printf("Comparing (%d,%d)[h=%d] with [ch=%d]\n", i, j, grid[i][j], height)
			if grid[i][j] > height {
				visible[i][j] = true
				height = grid[i][j]
			}
		}
	}

	// Compute going left
	for j := 1; j < len(inputs)-1; j++ {
		height := grid[0][j]
		for i := 1; i < len(inputs)-1; i++ {
			if grid[i][j] > height {
				visible[i][j] = true
				height = grid[i][j]
			}
		}
	}

	// Compute coming from the right
	for i := 1; i < len(inputs)-1; i++ {
		height := grid[i][len(inputs)-1]
		for j := len(inputs) - 2; j > 0; j-- {
			if grid[i][j] > height {
				visible[i][j] = true
				height = grid[i][j]
			}
		}
	}

	// Compute coming from above
	for j := len(inputs) - 2; j > 0; j-- {
		height := grid[len(inputs)-1][j]
		for i := len(inputs) - 2; i > 0; i-- {
			if grid[i][j] > height {
				visible[i][j] = true
				height = grid[i][j]
			}
		}
	}

	total := 0
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs); j++ {
			if visible[i][j] {
				total += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", total)

	score2 := 0

	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs); j++ {
			// Compute trees it can see above
			up, down, left, right := 0, 0, 0, 0

			// Compute up
			for a := i - 1; a >= 0; a-- {
				up += 1 // we see the tree next to us
				if grid[i][j] <= grid[a][j] {
					break
				}
			}

			// Compute left
			for a := j - 1; a >= 0; a-- {
				left += 1 // we see the tree next to us
				if grid[i][j] <= grid[i][a] {
					break
				}
			}

			// Compute right
			for a := j + 1; a < len(inputs); a++ {
				right += 1 // we see the tree next to us
				if grid[i][j] <= grid[i][a] {
					break
				}
			}

			// Compute down
			for a := i + 1; a < len(inputs); a++ {
				down += 1 // we see the tree next to us
				if grid[i][j] <= grid[a][j] {
					break
				}
			}

			temp := up * down * left * right
			if temp > score2 {
				score2 = temp
			}
		}
	}

	fmt.Printf("Part 2: %d\n", score2)
}
