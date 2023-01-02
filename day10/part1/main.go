package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := 0
	cpuX := 1
	signal := 20
	cycle := 1
	lines := bytes.Split(data, []byte("\n"))

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
			if cycle == signal {
				res += cpuX * cycle
				signal += 40
			}
			cycle++
			times--
		}
		cpuX += value
	}

	fmt.Println(res)
}
