package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make(map[string]bool, 0)
	numbers2 := make(map[string]bool, 0)
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

		// The map value is not really important
		numbers[code] = true
		numbers2[code] = true
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

	// Oxygen generator val
	oxygen := 0

	{
		for i := 0; i < length && len(numbers) > 1; i++ {

			// Compute the ones for what we have left
			ones := make(map[int]int, 0)

			for key := range numbers {
				for pos, val := range key {
					if val == '1' {
						ones[pos] += 1
					}
				}
			}

			char := '0'
			zeros := len(numbers) - ones[i]
			if ones[i] >= zeros {
				char = '1'
			}

			for key := range numbers {
				if rune(key[i]) != char {
					delete(numbers, key)
				}
			}
		}

		for key := range numbers {
			val, _ := strconv.ParseInt(key, 2, 64)
			oxygen = int(val)
		}
	}

	// Scrubber generator
	co2 := 0
	{
		for i := 0; i < length && len(numbers2) > 1; i++ {

			// Compute the ones for what we have left
			ones := make(map[int]int, 0)

			for key := range numbers2 {
				for pos, val := range key {
					if val == '1' {
						ones[pos] += 1
					}
				}
			}

			char := '1'
			zeros := len(numbers2) - ones[i]
			if zeros <= ones[i] {
				char = '0'
			}

			for key := range numbers2 {
				if rune(key[i]) != char {
					delete(numbers2, key)
				}
			}
		}

		for key := range numbers2 {
			val, _ := strconv.ParseInt(key, 2, 64)
			co2 = int(val)
		}
	}

	fmt.Printf("Part 2: %d\n", oxygen*co2)
}
