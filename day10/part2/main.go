package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := ""
	spriteX := 1
	endOfLine := 39
	cycle := 0
	lines := bytes.Split(data, []byte("\n"))
	buf := make([]byte, 40)

	for _, line := range lines {
		value, times := 0, 0
		command := line[:4]

		if string(command) == "noop" {
			value = 0
			times = 1
		} else {
			command = line[5:]
			value, _ = strconv.Atoi(string(command))
			times = 2
		}

		for times > 0 {
			// Draw pixel
			buf[cycle] = '.'
			// draw sprite
			if math.Abs(float64(spriteX-cycle)) < 2 {
				buf[cycle] = '#'
			}

			// update X value after draw
			if times == 1 {
				spriteX += value
			}

			// we are at the end of line
			if cycle == endOfLine {
				cycle = -1
				buf[endOfLine] = '\n'
				res += string(buf)
				buf = make([]byte, 40)
			}
			cycle++
			times--
		}
	}

	fmt.Println(res)
}
