package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	Rock = iota + 1
	Paper
	Scissors
)

const (
	lost = 0
	draw = 3
	won  = 6
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	points := make(map[string]int)
	points["X"] = Rock
	points["Y"] = Paper
	points["Z"] = Scissors

	scores := make(map[string]int)
	scores["A X"] = draw
	scores["A Y"] = won
	scores["A Z"] = lost
	scores["B X"] = lost
	scores["B Y"] = draw
	scores["B Z"] = won
	scores["C X"] = won
	scores["C Y"] = lost
	scores["C Z"] = draw

	result := 0

	for _, line := range bytes.Split(data, []byte("\n")) {
		result += scores[string(line)] + points[string(line[2])]
	}

	fmt.Println(result)
}
