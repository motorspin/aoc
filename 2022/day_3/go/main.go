package main

import (
	"bufio"
	"fmt"
	"os"
)

func computePoints(letter rune) int {
	if letter >= 'A' && letter <= 'Z' {
		return int(letter) - 'A' + 27
	}

	return int(letter) - 'a' + 1
}

func getCommonItem(compartment1 string, compartment2 string) rune {
	itemsIn1 := make(map[rune]bool)

	for i := 0; i < len(compartment1); i++ {
		itemsIn1[rune(compartment1[i])] = true
	}

	for i := 0; i < len(compartment2); i++ {
		_, ok := itemsIn1[rune(compartment2[i])]

		if ok {
			return rune(compartment2[i])
		}
	}

	fmt.Println("Something bad occurred")

	return '*'
}

func PerformP1(rucksack string) int {
	compartment1 := rucksack[0 : len(rucksack)/2]
	compartment2 := rucksack[len(rucksack)/2:]
	return computePoints(getCommonItem(compartment1, compartment2))
}

func PerformP2(rucksack1 string, rucksack2 string, rucksack3 string) int {
	// find common item across rucksacks
	items1 := make(map[rune]bool)
	items2 := make(map[rune]bool)

	for i := 0; i < len(rucksack1); i++ {
		items1[rune(rucksack1[i])] = true
	}

	for i := 0; i < len(rucksack2); i++ {
		items2[rune(rucksack2[i])] = true
	}

	var letter rune = '*'

	for i := 0; i < len(rucksack3); i++ {
		_, ok1 := items1[rune(rucksack3[i])]
		_, ok2 := items2[rune(rucksack3[i])]

		if ok1 && ok2 {
			letter = rune(rucksack3[i])
			break
		}
	}

	if letter == '*' {
		fmt.Println("Something bad happened")
	}

	return computePoints(letter)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	score := 0
	score2 := 0

	for {
		if !scanner.Scan() {
			break
		}

		rucksack1 := scanner.Text()
		score += PerformP1(rucksack1)

		if !scanner.Scan() {
			break
		}

		rucksack2 := scanner.Text()
		score += PerformP1(rucksack2)

		if !scanner.Scan() {
			break
		}

		rucksack3 := scanner.Text()
		score += PerformP1(rucksack3)

		score2 += PerformP2(rucksack1, rucksack2, rucksack3)
	}

	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Part 2: %d\n", score2)
}
