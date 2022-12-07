package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ones := make(map[int]int, 0)
	length := 0
	count := 0
	for scanner.Scan() {
		code := scanner.Text()
		length = len(code)
		count += 1

		for pos, char := range code {
			if char == '1' {
				ones[pos] += 1
			}
		}
	}

	gamma := ""
	epsilon := ""
	for i := 0; i < length; i++ {
		if ones[i] > count/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaI, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonI, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("Part 1: %d\n", gammaI*epsilonI)
}
