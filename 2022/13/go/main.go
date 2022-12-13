package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Data struct {
	isInteger bool // if true it's an integer, else it's a list (another Data)
	val       int
	children  []Data
}

func (d Data) String() string {
	if !d.isInteger {
		// Handle list
		val := "["

		for i := 0; i < len(d.children); i++ {
			if i == 0 {
				val += fmt.Sprintf("%s", d.children[i])
			} else {
				val += ", " + fmt.Sprintf("%s", d.children[i])
			}
		}

		return val + "]"
	}

	return fmt.Sprintf("%d", d.val)
}

type Packet struct {
	data []Data
}

func ParseList(text string) []Data {
	ret := make([]Data, 0)

	// Let's remove the surrounding square brackets
	trimmed := text[1 : len(text)-1]

	if len(trimmed) == 0 {
		return ret
	}

	elems := make([]string, 0)

	currentStart := 0
	bracketDepth := 0
	for i := 0; i < len(trimmed); i++ {
		if '[' == trimmed[i] {
			bracketDepth += 1
		}
		if ']' == trimmed[i] {
			bracketDepth -= 1
		}
		if ',' == trimmed[i] && bracketDepth == 0 {
			elems = append(elems, trimmed[currentStart:i])
			currentStart = i + 1
		}
	}

	elems = append(elems, trimmed[currentStart:])

	for _, elem := range elems {
		if elem[0] != '[' {
			// Not a list
			val, _ := strconv.ParseInt(elem, 10, 64)
			ret = append(ret, Data{true, int(val), nil})
		} else {
			ret = append(ret, Data{false, 0, ParseList(elem)})
		}
	}

	return ret
}

func NewPacket(text string) *Packet {
	return &Packet{ParseList(text)}
}

func (p Packet) String() string {
	return fmt.Sprintf("%s", p.data)
}

// 0 == no decision
// 1 == ordered
// 2 == not ordered
func isOrderedList(d1 Data, d2 Data) int {
	index := 0

	for {
		// Both lists have run out of items
		if index == len(d1.children) && index == len(d2.children) {
			return 0
		}

		// Left side ran out of items
		if index == len(d1.children) && index < len(d2.children) {
			return 1
		}

		// Right side ran out of items
		if index == len(d2.children) {
			return 2
		}

		// If items are integers
		if d1.children[index].isInteger && d2.children[index].isInteger {
			val := isOrderedInt(d1.children[index], d2.children[index])

			// If there was no decision, move on to the next iteration
			if val == 0 {
				index += 1
				continue
			}

			return val
		}

		// If items are lists
		if !d1.children[index].isInteger && !d2.children[index].isInteger {
			val := isOrderedList(d1.children[index], d2.children[index])

			if val == 0 {
				index += 1
				continue
			}

			return val
		}

		// If left is an integer and right is a list
		if d1.children[index].isInteger && !d2.children[index].isInteger {
			lhs := Data{false, 0, make([]Data, 1)}
			lhs.children[0] = Data{true, d1.children[index].val, nil}

			val := isOrderedList(lhs, d2.children[index])

			if val == 0 {
				index += 1
				continue
			}

			return val
		}

		// If left is a list and right is an integer
		if !d1.children[index].isInteger && d2.children[index].isInteger {
			rhs := Data{false, 0, make([]Data, 1)}
			rhs.children[0] = Data{true, d2.children[index].val, nil}

			val := isOrderedList(d1.children[index], rhs)

			if val == 0 {
				index += 1
				continue
			}

			return val
		}
	}
}

// 0 == no decision
// 1 == ordered
// 2 == not ordered
func isOrderedInt(d1 Data, d2 Data) int {
	if d1.val < d2.val {
		return 1
	} else if d1.val > d2.val {
		return 2
	}

	return 0
}

func IsOrdered(p1 *Packet, p2 *Packet) bool {
	lhs := Data{false, 0, p1.data}
	rhs := Data{false, 0, p2.data}

	val := isOrderedList(lhs, rhs)

	if val == 0 {
		panic("Equal!")
	}

	return val == 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum, count := 0, 0
	packets := make([]*Packet, 0)

	for scanner.Scan() {
		// Increment count as we are processing a new pair of packets
		count += 1

		// Grab the first packet
		packet1 := NewPacket(scanner.Text())
		packets = append(packets, packet1)

		// Grab the second packet
		scanner.Scan()
		packet2 := NewPacket(scanner.Text())
		packets = append(packets, packet2)

		if IsOrdered(packet1, packet2) {
			sum += count
		}

		// Scan the next line (will be empty)
		scanner.Scan()
	}

	fmt.Println("Part 1:", sum)

	// Insert divider packets
	packets = append(packets, NewPacket("[[2]]"))
	packets = append(packets, NewPacket("[[6]]"))

	// Sort the packets
	sort.Slice(packets, func(i int, j int) bool {
		return IsOrdered(packets[i], packets[j])
	})

	d1Index, d2Index := 0, 0

	for i, val := range packets {
		if "[[2]]" == fmt.Sprintf("%s", val) {
			d1Index = i + 1
		}

		if "[[6]]" == fmt.Sprintf("%s", val) {
			d2Index = i + 1
		}
	}

	decoderKey := d1Index * d2Index
	fmt.Println("Part 2:", decoderKey)
}
