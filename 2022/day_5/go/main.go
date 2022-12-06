package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Parse Stacks && Instructions
	stackStrings := make([]string, 0)
	instructions := make([]string, 0)
	parseStacks := true
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			parseStacks = false
			continue
		}

		if parseStacks {
			stackStrings = append(stackStrings, line)
		} else {
			instructions = append(instructions, line)
		}
	}

	// How many stacks do we have? Let's take a look at the last entry in stackStrings
	numStacks := len(strings.Fields(stackStrings[len(stackStrings)-1]))
	stacks := make([]Stack, numStacks)

	for i := len(stackStrings) - 1; i >= 0; i-- {
		//fmt.Println(stackStrings[i])
		if i == len(stackStrings)-1 {
			continue
		}

		for a := 0; a < len(stackStrings[i]); a++ {
			// A crate takes up 3 characters for '[' , the letter and, ']'
			// Each create is separated by a space
			// Spaces will be present for all 3 characters to indicate the lack of a create
			if stackStrings[i][a] != '[' && stackStrings[i][a] != ']' && stackStrings[i][a] != ' ' {
				//fmt.Println(string(stackStrings[i][a]), "at position", a)
				stacks[((a - 1) / 4)].Push(rune(stackStrings[i][a]))
			}
		}
	}

	//fmt.Println("Stacks: ", stackStrings)
	//fmt.Println("Instructions: ", instructions)

	for _, instruction := range instructions {
		var qty int
		var orig int
		var dest int
		fmt.Sscanf(instruction, "move %d from %d to %d", &qty, &orig, &dest)

		for i := 0; i < qty; i++ {
			val, _ := stacks[orig-1].Pop()
			stacks[dest-1].Push(val)
		}
		//fmt.Printf("Moving %d creates from stack %d to stack %d\n", qty, orig, dest)
	}

	fmt.Printf("Part 1: ")
	for _, stack := range stacks {
		char, _ := stack.Peek()
		fmt.Printf("%c", char)
	}
	fmt.Println("")
}
