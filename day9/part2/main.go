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

func (p *position) equal(p2 *position) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(data)

	result := make(map[position]struct{})
	knots := []*knot{}
	for i := 0; i < 10; i++ {
		knots = append(knots, &knot{&position{0, 0}})
	}

	head := &position{}
	tail := &position{}
	result[*tail] = struct{}{}

	for _, line := range strings.Split(lines, "\n") {
		dir, a := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		amount, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal(err)
		}

	outer:
		for i := 0; i < amount; i++ {
			switch dir {
			case "R":
				head.X++
			case "L":
				head.X--
			case "U":
				head.Y++
			case "D":
				head.Y--
			}

			for i, k := range knots {
				last := head
				if i > 0 {
					last = knots[i-1].pos
				}

				xDiff := float64(last.X - k.pos.X)
				yDiff := float64(last.Y - k.pos.Y)

				if math.Abs(xDiff) < 2 && math.Abs(yDiff) < 2 {
					continue
				}

				if xDiff > 0 {
					k.pos.X++
				} else if xDiff < 0 {
					k.pos.X--
				}
				if yDiff > 0 {
					k.pos.Y++
				} else if yDiff < 0 {
					k.pos.Y--
				}

				if i+1 == len(knots)-1 {
					if tail.equal(k.pos) {
						continue outer
					}
					result[*k.pos] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(result))
}

// 2471
