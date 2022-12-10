package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cycle := 1
	cycle2 := 0
	x := 1
	signalStrength := 0
	spriteStart := 1

	for scanner.Scan() {
		instruction := scanner.Text()

		if instruction == "noop" {
			cycle += 1
			cycle2 += 1

			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signalStrength += cycle * x
			}

			if cycle2%40 == spriteStart || cycle2%40 == spriteStart+1 || cycle2%40 == spriteStart+2 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			if cycle2 == 40 || cycle2 == 80 || cycle2 == 120 || cycle2 == 160 || cycle2 == 200 || cycle2 == 240 {
				fmt.Println()
			}
		} else {
			parts := strings.Split(instruction, " ")
			val, _ := strconv.ParseInt(parts[1], 10, 64)

			cycle += 1
			cycle2 += 1

			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signalStrength += cycle * x
			}

			if cycle2%40 == spriteStart || cycle2%40 == spriteStart+1 || cycle2%40 == spriteStart+2 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			if cycle2 == 40 || cycle2 == 80 || cycle2 == 120 || cycle2 == 160 || cycle2 == 200 || cycle2 == 240 {
				fmt.Println()
			}

			cycle += 1
			cycle2 += 1
			x += int(val)
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signalStrength += cycle * x
			}

			if cycle2%40 == spriteStart || cycle2%40 == spriteStart+1 || cycle2%40 == spriteStart+2 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			if cycle2 == 40 || cycle2 == 80 || cycle2 == 120 || cycle2 == 160 || cycle2 == 200 || cycle2 == 240 {
				fmt.Println()
			}

			spriteStart = x
		}
	}

	fmt.Println()
	fmt.Printf("Part 1: %d\n", signalStrength)
	fmt.Println("Part 2: See generated image")
}
