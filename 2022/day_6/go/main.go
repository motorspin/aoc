package main

import (
	"bufio"
	"fmt"
	"os"
)

type LastFour struct {
	s   []rune
	pos int
}

func NewLastFour() *LastFour {
	return &LastFour{make([]rune, 4), 0}
}

func (s *LastFour) Push(val rune) {
	if s.pos < 4 {
		s.s[s.pos] = val
	} else {
		s.s[0] = s.s[1]
		s.s[1] = s.s[2]
		s.s[2] = s.s[3]
		s.s[3] = val
	}
	s.pos += 1
}

func (s *LastFour) IsStart() bool {
	if s.pos < 4 {
		return false
	}

	if s.s[0] == s.s[1] || s.s[0] == s.s[2] || s.s[0] == s.s[3] {
		return false
	}

	if s.s[1] == s.s[2] || s.s[1] == s.s[3] {
		return false
	}

	if s.s[2] == s.s[3] {
		return false
	}

	return true
}

func (s *LastFour) Position() int {
	return s.pos
}

type LastFourteen struct {
	s   []rune
	pos int
	m   map[rune]int
}

func NewLastFourteen() *LastFourteen {
	return &LastFourteen{make([]rune, 14), 0, make(map[rune]int, 0)}
}

func (s *LastFourteen) Push(val rune) {
	if s.pos < 14 {
		s.s[s.pos] = val
	} else {
		s.m[s.s[0]] -= 1
		s.s = s.s[1:]
		s.s = append(s.s, val)
	}
	s.m[val] += 1
	s.pos += 1
}

func (s *LastFourteen) IsStart() bool {
	if s.pos < 14 {
		return false
	}

	for _, val := range s.m {
		if val > 1 {
			return false
		}
	}

	return true
}

func (s *LastFourteen) Position() int {
	return s.pos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	buffer := scanner.Text()
	lastFour := NewLastFour()

	fmt.Printf("Part 1: ")
	for i := 0; i < len(buffer); i++ {
		lastFour.Push(rune(buffer[i]))
		if lastFour.IsStart() {
			fmt.Println(lastFour.Position())
			break
		}
	}

	lastFourteen := NewLastFourteen()
	fmt.Printf("Part 2: ")
	for i := 0; i < len(buffer); i++ {
		lastFourteen.Push(rune(buffer[i]))
		if lastFourteen.IsStart() {
			fmt.Println(lastFourteen.Position())
			break
		}
	}
}
