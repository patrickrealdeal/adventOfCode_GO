// Day 3
package main

import (
	"adventOfCode/day3/intset"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	alphabet := []byte{}
	points := make(map[byte]int)
	appeard := intset.IntSet{}

	for i, n := 'A', 26; i <= 'Z'; i, n = i+1, n+1 {
		alphabet = append(alphabet, byte(i))
		points[byte(i)] = n + 1
	}

	for i, n := 'a', 0; i <= 'z'; i, n = i+1, n+1 {
		alphabet = append(alphabet, byte(i))
		points[byte(i)] = n + 1
	}

	res := 0

	for _, line := range bytes.Split(data, []byte("\n")) {
		left, right := line[:len(line)/2], line[len(line)/2:]

		for _, c := range line {
			appeard.Add(int(c))
		}

		for _, value := range appeard.Values() {
			if bytes.ContainsRune(left, rune(value)) && bytes.ContainsRune(right, rune(value)) {
				res += points[byte(value)]
			}
		}
		appeard.Clear()
	}
	fmt.Println(res)
}

// 7763
