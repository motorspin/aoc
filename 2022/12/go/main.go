package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coord struct {
	i int
	j int
}

type CoordWithInfo struct {
	Coord
	steps   int
	visited *map[Coord]bool
}

type Queue struct {
	items []CoordWithInfo
}

func NewQueue() *Queue {
	return &Queue{make([]CoordWithInfo, 0)}
}

func (q *Queue) Push(coord CoordWithInfo) {
	q.items = append(q.items, coord)
}

func (q *Queue) Pop() (coord CoordWithInfo) {
	coord = q.items[0]
	q.items = q.items[1:]
	return
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func solve(hmap [][]rune, part int) int {
	possibleStarts := make([]Coord, 0)
	end := Coord{0, 0}

	for i := 0; i < len(hmap); i++ {
		for j := 0; j < len(hmap[0]); j++ {
			if hmap[i][j] == 'S' {
				hmap[i][j] = 'a'
				possibleStarts = append(possibleStarts, Coord{i, j})
			}

			if part == 2 {
				if hmap[i][j] == 'a' {
					possibleStarts = append(possibleStarts, Coord{i, j})
				}
			}

			if hmap[i][j] == 'E' {
				end = Coord{i, j}
				hmap[i][j] = 'z'
			}
		}
	}

	minSteps := math.MaxInt
	queue := NewQueue()

	// We should be able to share the same visited map since this is a BFS
	visitedMap := make(map[Coord]bool, 0)

	for _, val := range possibleStarts {
		queue.Push(CoordWithInfo{val, 0, &visitedMap})
	}

	for !queue.IsEmpty() {
		coord := queue.Pop()
		if (*coord.visited)[coord.Coord] {
			continue
		}
		(*coord.visited)[coord.Coord] = true

		if coord.i == end.i && coord.j == end.j {
			// We reached the end
			if coord.steps < minSteps {
				minSteps = coord.steps
			}

			// We're at the end, nothing left to do for this possible search
			continue
		}

		// Let's see if we can go up (if we aren't going out of bounds, and our
		// current letter+1 is >= square above us)
		if coord.i != 0 && hmap[coord.i][coord.j]+1 >= hmap[coord.i-1][coord.j] {
			var newCoord CoordWithInfo
			newCoord.i = coord.i - 1
			newCoord.j = coord.j
			newCoord.steps = coord.steps + 1
			newCoord.visited = coord.visited
			queue.Push(newCoord)
		}

		// Let's see if we can go left (if we aren't going out of bounds, and our
		// current letter+1 is >= square to the left of us)
		if coord.j != 0 && hmap[coord.i][coord.j]+1 >= hmap[coord.i][coord.j-1] {
			var newCoord CoordWithInfo
			newCoord.i = coord.i
			newCoord.j = coord.j - 1
			newCoord.steps = coord.steps + 1
			newCoord.visited = coord.visited
			queue.Push(newCoord)
		}

		// Let's see if we can go right (if we aren't going out of bounds, and our
		// current letter+1 is >= square to the right of us)
		if coord.j != len(hmap[0])-1 && hmap[coord.i][coord.j]+1 >= hmap[coord.i][coord.j+1] {
			var newCoord CoordWithInfo
			newCoord.i = coord.i
			newCoord.j = coord.j + 1
			newCoord.steps = coord.steps + 1
			newCoord.visited = coord.visited
			queue.Push(newCoord)
		}

		// Let's see if we can go down (if we aren't going out of bounds, and our
		// current letter+1 is >= square below us)
		//fmt.Println(coord.i, len(hmap)-1, hmap[coord.i][coord.j]+1, hmap[coord.i+1][coord.j])
		if (coord.i != len(hmap)-1) && ((hmap[coord.i][coord.j] + 1) >= hmap[coord.i+1][coord.j]) {
			var newCoord CoordWithInfo
			newCoord.i = coord.i + 1
			newCoord.j = coord.j
			newCoord.steps = coord.steps + 1
			newCoord.visited = coord.visited
			queue.Push(newCoord)
		}
	}

	return minSteps
}

func main() {
	hmap := make([][]rune, 0)
	hmap2 := make([][]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		hrow := make([]rune, 0)
		hrow2 := make([]rune, 0)
		for _, val := range scanner.Text() {
			hrow = append(hrow, val)
			hrow2 = append(hrow2, val)
		}

		hmap = append(hmap, hrow)
		hmap2 = append(hmap2, hrow2)
	}

	fmt.Println("Part 1:", solve(hmap, 1))
	fmt.Println("Part 2:", solve(hmap2, 2))
}
