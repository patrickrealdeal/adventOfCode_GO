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
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	alphabet := make([]byte, 52)
	points := make(map[byte]int, 52)
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
	block := [][]byte{}

	for i, line := range bytes.Split(data, []byte("\n")) {
		block = append(block, line)

		if i%3 == 2 {
			for _, elem := range block {
				for _, c := range elem {
					appeard.Add(int(c))
				}
			}

			for _, letter := range appeard.Values() {
				if bytes.ContainsRune(block[0], rune(letter)) &&
					bytes.ContainsRune(block[1], rune(letter)) &&
					bytes.ContainsRune(block[2], rune(letter)) {
					res += points[letter]
				}
			}
			block = [][]byte{}
			appeard.Clear()
		}
	}
	fmt.Println(res)
}

// 2569
