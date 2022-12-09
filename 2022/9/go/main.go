package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	i int
	j int
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}

func printGrid(grid [][]bool, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}

		fmt.Println()
	}
}

//func isOverlapped(headi int, headj int, taili int, tailj int) bool {
//	return headi == taili && headj == tailj
//}

func isAdjacent(headi int, headj int, taili int, tailj int) bool {

	rowDifference := abs(headi - taili)
	colDifference := abs(headj - tailj)

	return rowDifference <= 1 && colDifference <= 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// This feels hacky... but we'll do it anyway
	size := 2000
	grid := make([][]bool, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]bool, size)
	}

	starti, startj := size/2, size/2
	headi, headj := starti, startj
	taili, tailj := starti, startj
	grid[taili][tailj] = true
	instructions := make([]string, 0)

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		val, _ := strconv.ParseInt(parts[1], 10, 64)
		amount := int(val)

		// Let's assume the T is in a position where it's next to the H when we start
		for i := 0; i < amount; i++ {
			if parts[0] == "R" {
				headj += 1

				if !isAdjacent(headi, headj, taili, tailj) {
					taili = headi
					tailj = headj - 1
				}
			} else if parts[0] == "L" {
				headj -= 1

				if !isAdjacent(headi, headj, taili, tailj) {
					taili = headi
					tailj = headj + 1
				}
			} else if parts[0] == "U" {
				headi -= 1

				if !isAdjacent(headi, headj, taili, tailj) {
					taili = headi + 1
					tailj = headj
				}
			} else {
				// down
				headi += 1

				if !isAdjacent(headi, headj, taili, tailj) {
					taili = headi - 1
					tailj = headj
				}
			}
			grid[taili][tailj] = true
		}
	}

	count := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] {
				count += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count)

	// Need to make the soln above a bit more generic
	grid2 := make([][]bool, size)

	for i := 0; i < size; i++ {
		grid2[i] = make([]bool, size)
	}

	knots := make([]Coord, 0)
	knots = append(knots, Coord{size / 2, size / 2}) // head
	knots = append(knots, Coord{size / 2, size / 2}) // 1
	knots = append(knots, Coord{size / 2, size / 2}) // 2
	knots = append(knots, Coord{size / 2, size / 2}) // 3
	knots = append(knots, Coord{size / 2, size / 2}) // 4
	knots = append(knots, Coord{size / 2, size / 2}) // 5
	knots = append(knots, Coord{size / 2, size / 2}) // 6
	knots = append(knots, Coord{size / 2, size / 2}) // 7
	knots = append(knots, Coord{size / 2, size / 2}) // 8
	knots = append(knots, Coord{size / 2, size / 2}) // 9

	// Every knot is starting in the same place, so we can mark this as visited
	grid2[knots[0].i][knots[0].j] = true

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		val, _ := strconv.ParseInt(parts[1], 10, 64)
		amount := int(val)

		// Let's assume the T is in a position where it's next to the H when we start
		for i := 0; i < amount; i++ {
			if parts[0] == "R" {
				knots[0].j += 1
			} else if parts[0] == "L" {
				knots[0].j -= 1
			} else if parts[0] == "U" {
				knots[0].i -= 1
			} else {
				// down
				knots[0].i += 1
			}

			// Go through all the knots one-by-one
			for a := 1; a < len(knots); a++ {
				if !isAdjacent(knots[a-1].i, knots[a-1].j, knots[a].i, knots[a].j) {
					colDifference := knots[a].j - knots[a-1].j
					rowDifference := knots[a].i - knots[a-1].i

					if colDifference == 0 {
						if knots[a-1].i > knots[a].i {
							knots[a].i = knots[a-1].i - 1
						} else {
							knots[a].i = knots[a-1].i + 1
						}
					} else if rowDifference == 0 {
						if knots[a-1].j > knots[a].j {
							knots[a].j = knots[a-1].j - 1
						} else {
							knots[a].j = knots[a-1].j + 1
						}
					} else {
						// Need to move diagonally
						if knots[a-1].i > knots[a].i {
							knots[a].i += 1
						} else {
							knots[a].i -= 1
						}

						if knots[a-1].j > knots[a].j {
							knots[a].j += 1
						} else {
							knots[a].j -= 1
						}
					}
				}
			}
			grid2[knots[len(knots)-1].i][knots[len(knots)-1].j] = true
		}
	}

	count2 := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid2[i][j] {
				count2 += 1
			}
		}
	}

	fmt.Printf("Part 2: %d\n", count2)
}
