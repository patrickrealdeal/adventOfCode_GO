package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type stack struct {
	cols  map[int][]string
	moves map[int][]int
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := string(data)

	s := &stack{
		cols:  make(map[int][]string, 9),
		moves: make(map[int][]int, 502),
	}
	m := strings.Index(lines, "m")
	movs := lines[m:]
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	for i, line := range strings.Split(movs, "\n") {
		submatchall := re.FindAllString(line, -1)
		for _, c := range submatchall {
			parsed, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			s.moves[i] = append(s.moves[i], parsed)
		}
	}

	boxes := lines[:m]
	boxes = strings.Trim(boxes, "\n")

	for _, line := range strings.Split(boxes, "\n") {
		for j, c := range line {
			if unicode.IsUpper(c) {
				s.cols[j/4+1] = append(s.cols[j/4+1], string(c))
			}
		}
	}

	for i := 0; i < len(s.cols)+1; i++ {
		reverse(s.cols[i])
	}

	for i := 0; i < len(s.moves); i++ {
		move(s.moves[i][0], s.moves[i][1], s.moves[i][2], s)
	}

	res := ""
	for i := 1; i < len(s.cols)+1; i++ {
		res += s.cols[i][len(s.cols[i])-1]
	}

	fmt.Println(res)
}

func move(n, from, to int, s *stack) {
	moving := s.cols[from][len(s.cols[from])-n:]
	s.cols[to] = append(s.cols[to], moving...)
	if len(s.cols[from]) > 0 {
		s.cols[from] = s.cols[from][:len(s.cols[from])-n]
	}
}

func reverse[S []E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// RWLWGJGFD
