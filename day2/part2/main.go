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
	points["A"] = Rock
	points["B"] = Paper
	points["C"] = Scissors

	points["X"] = Rock
	points["Y"] = Paper
	points["Z"] = Scissors

	result := 0
	for _, line := range bytes.Split(data, []byte("\n")) {
		ln := bytes.Split(line, []byte(" "))
		oppMove, myMove := string(ln[0]), string(ln[1])

		if myMove == "X" {
			if oppMove == "A" {
				result += lost + Scissors
			}
			if oppMove == "B" {
				result += lost + Rock
			}
			if oppMove == "C" {
				result += lost + Paper
			}
			continue
		}

		if myMove == "Y" {
			result += draw + points[oppMove]
			continue
		}

		if myMove == "Z" {
			if oppMove == "A" {
				result += won + Paper
			}
			if oppMove == "B" {
				result += won + Scissors
			}
			if oppMove == "C" {
				result += won + Rock
			}
			continue
		}
	}

	fmt.Println(result)
}

// 11186
