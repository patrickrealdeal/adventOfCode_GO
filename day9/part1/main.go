package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	X int
	Y int
}

type knot struct {
	pos *position
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(data)

	res := make(map[position]int)

	head := knot{&position{0, 0}}
	tail := knot{&position{0, 0}}

	res[*tail.pos]++

	for _, line := range strings.Split(lines, "\n") {
		dir, a := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		amount, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < amount; i++ {
			switch dir {
			case "U":
				head.pos.Y += 1
			case "D":
				head.pos.Y -= 1
			case "R":
				head.pos.X += 1
			case "L":
				head.pos.X -= 1
			}

			xDiff := float64(head.pos.X - tail.pos.X)
			yDiff := float64(head.pos.Y - tail.pos.Y)
			if math.Abs(xDiff) < 2 && math.Abs(yDiff) < 2 {
				continue
			}

			if xDiff > 0 {
				tail.pos.X++
			} else if xDiff < 0 {
				tail.pos.X--
			}
			if yDiff > 0 {
				tail.pos.Y++
			} else if yDiff < 0 {
				tail.pos.Y--
			}
			res[*tail.pos]++
		}
	}

	fmt.Println(len(res))
}

// 6067
