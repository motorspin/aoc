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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	buffer := scanner.Text()
	lastFour := NewLastFour()

	for i := 0; i < len(buffer); i++ {
		lastFour.Push(rune(buffer[i]))
		if lastFour.IsStart() {
			fmt.Println(lastFour.Position())
			break
		}
	}
}
