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

	for _, line := range bytes.Split(data, []byte("\n")) {
		comma := bytes.Split(line, []byte(","))
		left, right := bytes.Split(comma[0], []byte("-")), bytes.Split(comma[1], []byte("-"))
		l1, l2 := toInt(left[0]), toInt(left[1])
		r1, r2 := toInt(right[0]), toInt(right[1])

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

// 849
