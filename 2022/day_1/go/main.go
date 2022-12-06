package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	elves := make([]int, 0)

	var total int = 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()

		if len(word) > 0 {
			val, err := strconv.ParseInt(word, 10, 32)
			if err != nil {
				panic(err)
			}
			total = total + int(val)
		} else {
			elves = append(elves, total)
			total = 0
		}
	}

	if total > 0 {
		elves = append(elves, total)
	} else {
		fmt.Println("No data!")
		return
	}

	sort.Ints(elves)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Printf("Part 1: %d\n", elves[0])
	fmt.Printf("Part 2: %d\n", elves[0]+elves[1]+elves[2])
}
