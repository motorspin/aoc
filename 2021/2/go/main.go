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
	depth := 0
	forward := 0

	depth2 := 0
	forward2 := 0
	aim2 := 0
	for scanner.Scan() {
		instruction := scanner.Text()
		parts := strings.Split(instruction, " ")
		val, _ := strconv.ParseInt(parts[1], 10, 64)

		// Part 1
		if parts[0] == "forward" {
			forward += int(val)
		} else if parts[0] == "down" {
			depth += int(val)
		} else if parts[0] == "up" {
			depth -= int(val)
		}

		// Part 2
		if parts[0] == "forward" {
			forward2 += int(val)
			depth2 += int(val) * aim2
		} else if parts[0] == "down" {
			aim2 += int(val)
		} else if parts[0] == "up" {
			aim2 -= int(val)
		}
	}

	fmt.Printf("Part 1: %d\n", depth*forward)
	fmt.Printf("Part 2: %d\n", depth2*forward2)
}
