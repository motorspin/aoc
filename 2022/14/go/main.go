package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cell int64

const (
	Air Cell = iota
	Rock
	Sand
)

type Coord struct {
	x int
	y int
}

func NewCoordFromText(text string) Coord {
	parts := strings.Split(text, ",")
	x, _ := strconv.ParseInt(parts[0], 10, 64)
	y, _ := strconv.ParseInt(parts[1], 10, 64)
	return Coord{int(x), int(y)}
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func solve(cave map[Coord]Cell, bottom int, part int) int {
	floor := bottom + 2
	source := Coord{500, 0}

mainloop:
	for {
		// Produce a unit of sand
		sand := source

		for {
			// Are we already going into the void?
			if part == 1 && sand.y > bottom {
				break mainloop
			}

			// The floor checks are only needed for part 2 but they don't harm part 1
			below := Coord{sand.x, sand.y + 1}
			if cave[below] == Air && below.y != floor {
				sand = below
				continue
			}

			left := Coord{sand.x - 1, sand.y + 1}
			if cave[left] == Air && left.y != floor {
				sand = left
				continue
			}

			right := Coord{sand.x + 1, sand.y + 1}
			if cave[right] == Air && right.y != floor {
				sand = right
				continue
			}

			// We are blocked if we have gotten here, deposit sand
			cave[sand] = Sand

			if part == 2 && sand == source {
				break mainloop
			}

			break
		}
	}

	count := 0
	for _, val := range cave {
		if val == Sand {
			count += 1
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cave := make(map[Coord]Cell, 0)
	bottom := math.MinInt

	for scanner.Scan() {
		path := scanner.Text()
		lines := strings.Split(path, " -> ")

		for fromIndex, toIndex := 0, 1; toIndex < len(lines); fromIndex, toIndex = fromIndex+1, toIndex+1 {
			from := NewCoordFromText(lines[fromIndex])
			to := NewCoordFromText(lines[toIndex])

			// Draw horizontal lines
			for i := min(from.x, to.x); i <= max(from.x, to.x); i++ {
				cave[Coord{i, from.y}] = Rock
			}

			// Draw vertical lines
			for i := min(from.y, to.y); i <= max(from.y, to.y); i++ {
				cave[Coord{from.x, i}] = Rock
			}

			if bottom < max(from.y, to.y) {
				bottom = max(from.y, to.y)
			}
		}
	}

	fmt.Println("Part 1:", solve(cave, bottom, 1))
	fmt.Println("Part 2:", solve(cave, bottom, 2))
}
