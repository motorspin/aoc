package main

import (
	"bufio"
	"fmt"
	"os"
)

// A or X: rock
// B or Y: paper
// C or Z: scissors

func computeScore(playerOne string, playerTwo string) int {
	score := 0

	if "X" == playerOne {
		score += 1 // for playing rock

		if "A" == playerTwo {
			score += 3
		} else if "B" == playerTwo {
			score += 0
		} else {
			score += 6
		}

	} else if "Y" == playerOne {
		score += 2 // for playing paper

		if "A" == playerTwo {
			score += 6
		} else if "B" == playerTwo {
			score += 3
		} else {
			score += 0
		}

	} else {
		score += 3 // for playing scissors

		if "A" == playerTwo {
			score += 0
		} else if "B" == playerTwo {
			score += 6
		} else {
			score += 3
		}
	}

	return score
}

func computeScorePart2(opponent string, code string) int {
	score := 0

	if "A" == opponent {
		// they have rock

		if "X" == code {
			// we need to lose, pick scissors
			score += 3 + 0
		} else if "Y" == code {
			// we need to tie, pick rock
			score += 1 + 3
		} else if "Z" == code {
			// we need to win, pick paper
			score += 2 + 6
		}
	} else if "B" == opponent {
		// they have paper

		if "X" == code {
			// we need to lose, pick rock
			score += 1 + 0
		} else if "Y" == code {
			// we need to tie, pick paper
			score += 2 + 3
		} else {
			// we need to win, pick scissors
			score += 3 + 6
		}
	} else {
		// they have scissors

		if "X" == code {
			// we need to lose, pick paper
			score += 2 + 0
		} else if "Y" == code {
			// we need to tie, pick scissors
			score += 3 + 3
		} else {
			// we need to win, pick rock
			score += 1 + 6
		}
	}

	return score
}

func main() {
	fmt.Println(os.Args)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	score := 0
	scorePart2 := 0
	for {
		if !scanner.Scan() {
			break
		}
		playerOne := scanner.Text()

		if !scanner.Scan() {
			panic("Missing part of pair")
		}
		playerTwo := scanner.Text()

		score += computeScore(playerTwo, playerOne)
		scorePart2 += computeScorePart2(playerOne, playerTwo)
	}

	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Part 2: %d\n", scorePart2)
}
