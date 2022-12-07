package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convert(blah string) int {
	val, err := strconv.ParseInt(blah, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(val)
}

func isRangeContained(start int, end int, beg int, last int) bool {
	return beg >= start && last <= end
}

func doesRangeOverlap2(start int, end int, beg int, last int) bool {
	return start <= last && end >= beg
}

func doesRangeOverlap(start int, end int, beg int, last int) bool {
	for i := start; i <= end; i++ {
		for x := beg; x <= last; x++ {
			if i == x {
				return true
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	count := 0
	overlaps := 0

	for scanner.Scan() {
		pair := scanner.Text()
		assignments := strings.Split(pair, ",")
		assignment1 := strings.Split(assignments[0], "-")
		assignment2 := strings.Split(assignments[1], "-")
		//fmt.Println(assignments)
		//fmt.Println(assignment1)
		//fmt.Println(assignment2)
		//fmt.Printf("Assigment1 contained in Assigment2: %t\n", isRangeContained(convert(assignment2[0]), convert(assignment2[1]), convert(assignment1[0]), convert(assignment1[1])))
		//fmt.Printf("Assigment2 contained in Assigment1: %t\n", isRangeContained(convert(assignment1[0]), convert(assignment1[1]), convert(assignment2[0]), convert(assignment2[1])))
		//fmt.Println("")

		if doesRangeOverlap2(convert(assignment2[0]), convert(assignment2[1]), convert(assignment1[0]), convert(assignment1[1])) {
			overlaps += 1
		}

		if assignments[0] == assignments[1] {
			//fmt.Println(assignments)
			count += 1
			continue
		}

		if isRangeContained(convert(assignment2[0]), convert(assignment2[1]), convert(assignment1[0]), convert(assignment1[1])) {
			count += 1
		}

		if isRangeContained(convert(assignment1[0]), convert(assignment1[1]), convert(assignment2[0]), convert(assignment2[1])) {
			count += 1
		}
	}

	fmt.Printf("Part 1: %d\n", count)
	fmt.Printf("Part 2: %d\n", overlaps)
}
