package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := 0

	for _, line := range bytes.Split(data, []byte("\n")) {
		l1, l2, r1, r2 := parseLine(line)

		if l2 >= r1 && l1 <= r2 {
			res += 1
		}
	}

	fmt.Println(res)
}

func toInt(b []byte) int {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return i
}

func parseLine(line []byte) (a, b, c, d int) {
	comma := bytes.Split(line, []byte(","))
	left, right := bytes.Split(comma[0], []byte("-")), bytes.Split(comma[1], []byte("-"))
	a, b = toInt(left[0]), toInt(left[1])
	c, d = toInt(right[0]), toInt(right[1])
	return
}

// 849
