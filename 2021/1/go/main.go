package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ints := make([]int, 0)

	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		ints = append(ints, int(val))
	}

	count := 0
	last := ints[0]

	for i := 1; i < len(ints); i++ {
		if ints[i] > last {
			count++
		}
		last = ints[i]
	}

	fmt.Printf("Part 1: %d\n", count)

	count = 0
	last = ints[0] + ints[1] + ints[2]

	for i := 3; i < len(ints); i++ {
		sum := ints[i] + ints[i-1] + ints[i-2]
		if sum > last {
			count++
		}
		last = sum
	}

	fmt.Printf("Part 2: %d\n", count)
}
