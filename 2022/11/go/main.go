package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items          []int
	op             func(int) int
	test           int // this is the divisibility value
	fDest          int // where to throw if test is false
	tDest          int // where to throw if test is true
	itemsInspected int
}

func (m Monkey) String() string {
	return fmt.Sprintf("Items %s\nTest: divisible by %d\n\tIf true: throw to monkey %d\n\tIf false: throw to monkey: %d\nItems Inspected: %d\n", fmt.Sprint(m.items), m.test, m.tDest, m.fDest, m.itemsInspected)
}

func main() {
	monkeys := make([]Monkey, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var monkey Monkey

		// We don't do anything with the first Scanned line, it's just "Monkey
		// %d", go to next for items

		// Get Starting Items
		scanner.Scan()

		itemsText := scanner.Text()
		parts := strings.Split(itemsText, ":")
		// We don't care about parts[0], just parts[1]
		itemStrings := strings.Split(parts[1], ",")

		for _, v := range itemStrings {
			val, _ := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
			monkey.items = append(monkey.items, int(val))
		}

		// Parse Operation
		scanner.Scan()

		opHalves := strings.Split(scanner.Text(), "=")
		opParts := strings.Split(strings.TrimSpace(opHalves[1]), " ")

		monkey.op = func(old int) int {
			lhs := 0
			rhs := 0

			if opParts[0] == "old" {
				lhs = old
			} else {
				val, _ := strconv.ParseInt(opParts[0], 10, 64)
				lhs = int(val)
			}

			if opParts[2] == "old" {
				rhs = old
			} else {
				val, _ := strconv.ParseInt(opParts[2], 10, 64)
				rhs = int(val)
			}

			if opParts[1] == "+" {
				return lhs + rhs
			} else if opParts[1] == "*" {
				return lhs * rhs
			} else if opParts[1] == "/" {
				return lhs / rhs
			}

			return lhs - rhs
		}

		// Parse Test
		scanner.Scan()
		fmt.Sscanf(strings.TrimSpace(scanner.Text()), "Test: divisible by %d", &monkey.test)

		// Parse True Destination
		scanner.Scan()
		fmt.Sscanf(strings.TrimSpace(scanner.Text()), "If true: throw to monkey %d", &monkey.tDest)

		// Parse False Destination
		scanner.Scan()
		fmt.Sscanf(strings.TrimSpace(scanner.Text()), "If false: throw to monkey %d", &monkey.fDest)

		// Parse Empty Line After Monkey Processing
		scanner.Scan()

		monkeys = append(monkeys, monkey)
	}

	monkeys2 := make([]Monkey, len(monkeys))
	for i, monkey := range monkeys {
		monkeys2[i].items = monkey.items
		monkeys2[i].op = monkey.op
		monkeys2[i].test = monkey.test
		monkeys2[i].fDest = monkey.fDest
		monkeys2[i].tDest = monkey.tDest

	}

	// Perform simulation (part 1)
	for round := 0; round < 20; round++ {

		for n := 0; n < len(monkeys); n++ {
			for len(monkeys[n].items) > 0 {
				// Pull an item off the queue
				item := monkeys[n].items[0]
				monkeys[n].items = monkeys[n].items[1:]

				item = monkeys[n].op(item) / 3

				if item%monkeys[n].test != 0 {
					monkeys[monkeys[n].fDest].items = append(monkeys[monkeys[n].fDest].items, item)
				} else {
					monkeys[monkeys[n].tDest].items = append(monkeys[monkeys[n].tDest].items, item)
				}

				// Increase monkey item inspection count
				monkeys[n].itemsInspected += 1
			}
		}
	}

	inspectedItems := make([]int, len(monkeys))
	for n, monkey := range monkeys {
		inspectedItems[n] = monkey.itemsInspected
	}

	sort.Ints(inspectedItems)
	fmt.Printf("Part 1: %d\n", inspectedItems[len(inspectedItems)-1]*inspectedItems[len(inspectedItems)-2])

	// common divisor
	cd := 1

	for _, monkey := range monkeys {
		cd *= monkey.test
	}

	// Perform simulation (part 2)
	for round := 0; round < 10_000; round++ {
		for n := 0; n < len(monkeys2); n++ {
			for len(monkeys2[n].items) > 0 {
				// Pull an item off the queue
				item := monkeys2[n].items[0]
				monkeys2[n].items = monkeys2[n].items[1:]

				item = monkeys2[n].op(item) % int(cd)

				if item%monkeys2[n].test != 0 {
					monkeys2[monkeys2[n].fDest].items = append(monkeys2[monkeys2[n].fDest].items, item)
				} else {
					monkeys2[monkeys2[n].tDest].items = append(monkeys2[monkeys2[n].tDest].items, item)
				}

				// Increase monkey item inspection count
				monkeys2[n].itemsInspected += 1
			}
		}
	}

	inspectedItems2 := make([]int, len(monkeys2))
	for n, monkey := range monkeys2 {
		inspectedItems2[n] = monkey.itemsInspected
	}

	sort.Ints(inspectedItems2)
	fmt.Printf("Part 2: %d\n", inspectedItems2[len(inspectedItems2)-1]*inspectedItems2[len(inspectedItems2)-2])
}
