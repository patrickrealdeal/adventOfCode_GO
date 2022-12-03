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

	oppmoves := make(map[int]string)
	oppmoves[Rock] = "A"
	oppmoves[Paper] = "B"
	oppmoves[Scissors] = "C"
	mymoves := make(map[int]string)
	mymoves[Rock] = "X"
	mymoves[Paper] = "Y"
	mymoves[Scissors] = "Z"

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

		if myMove == mymoves[Rock] {
			if oppMove == oppmoves[Rock] {
				result += lost + Scissors
			}
			if oppMove == oppmoves[Paper] {
				result += lost + Rock
			}
			if oppMove == oppmoves[Scissors] {
				result += lost + Paper
			}
			continue
		}

		if myMove == mymoves[Paper] {
			result += draw + points[oppMove]
			continue
		}

		if myMove == mymoves[Scissors] {
			if oppMove == oppmoves[Rock] {
				result += won + Paper
			}
			if oppMove == oppmoves[Paper] {
				result += won + Scissors
			}
			if oppMove == oppmoves[Scissors] {
				result += won + Rock
			}
			continue
		}
	}

	fmt.Println(result)
}

// 11186
