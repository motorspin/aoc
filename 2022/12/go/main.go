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

type CoordWithSteps struct {
	Coord
	steps int
}

type Queue struct {
	items []CoordWithSteps
}

func NewQueue() *Queue {
	return &Queue{make([]CoordWithSteps, 0)}
}

func (q *Queue) Push(coord CoordWithSteps) {
	q.items = append(q.items, coord)
}

func (q *Queue) Pop() (coord CoordWithSteps) {
	coord = q.items[0]
	q.items = q.items[1:]
	return
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	hmap := make([][]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)
	start := Coord{0, 0}
	end := Coord{0, 0}

	for scanner.Scan() {
		hrow := make([]rune, 0)
		for _, val := range scanner.Text() {
			hrow = append(hrow, val)
		}

		hmap = append(hmap, hrow)
	}

	possibleStarts := make([]Coord, 0)

	for i := 0; i < len(hmap); i++ {
		for j := 0; j < len(hmap[0]); j++ {
			if hmap[i][j] == 'S' {
				start = Coord{i, j}
				hmap[i][j] = 'a'
				possibleStarts = append(possibleStarts, start)
			}

			if hmap[i][j] == 'a' {
				possibleStarts = append(possibleStarts, Coord{i, j})
			}

			if hmap[i][j] == 'E' {
				end = Coord{i, j}
				hmap[i][j] = 'z'
			}
		}
	}

	minSteps := math.MaxInt

	queue := NewQueue()
	visited := make(map[Coord]bool, 0)

	for _, val := range possibleStarts {
		queue.Push(CoordWithSteps{val, 0})
	}

	for !queue.IsEmpty() {
		coord := queue.Pop()
		if visited[coord.Coord] {
			continue
		}
		visited[coord.Coord] = true

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
		if coord.i != 0 && hmap[coord.i][coord.j]+1 >= hmap[coord.i-1][coord.j] && !visited[Coord{coord.i - 1, coord.j}] {
			var newCoord CoordWithSteps
			newCoord.i = coord.i - 1
			newCoord.j = coord.j
			newCoord.steps = coord.steps + 1
			queue.Push(newCoord)
		}

		// Let's see if we can go left (if we aren't going out of bounds, and our
		// current letter+1 is >= square to the left of us)
		if coord.j != 0 && hmap[coord.i][coord.j]+1 >= hmap[coord.i][coord.j-1] && !visited[Coord{coord.i, coord.j - 1}] {
			var newCoord CoordWithSteps
			newCoord.i = coord.i
			newCoord.j = coord.j - 1
			newCoord.steps = coord.steps + 1
			queue.Push(newCoord)
		}

		// Let's see if we can go right (if we aren't going out of bounds, and our
		// current letter+1 is >= square to the right of us)
		if coord.j != len(hmap[0])-1 && hmap[coord.i][coord.j]+1 >= hmap[coord.i][coord.j+1] && !visited[Coord{coord.i, coord.j + 1}] {
			var newCoord CoordWithSteps
			newCoord.i = coord.i
			newCoord.j = coord.j + 1
			newCoord.steps = coord.steps + 1
			queue.Push(newCoord)
		}

		// Let's see if we can go down (if we aren't going out of bounds, and our
		// current letter+1 is >= square below us)
		//fmt.Println(coord.i, len(hmap)-1, hmap[coord.i][coord.j]+1, hmap[coord.i+1][coord.j])
		if (coord.i != len(hmap)-1) && ((hmap[coord.i][coord.j] + 1) >= hmap[coord.i+1][coord.j]) && !visited[Coord{coord.i + 1, coord.j}] {
			var newCoord CoordWithSteps
			newCoord.i = coord.i + 1
			newCoord.j = coord.j
			newCoord.steps = coord.steps + 1
			queue.Push(newCoord)
		}
	}

	fmt.Println("Part 1:", minSteps)
}
