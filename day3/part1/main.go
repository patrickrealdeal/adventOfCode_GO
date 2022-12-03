// Day 3
package main

import (
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
	appeard := make(map[byte]int)

	for i, n := 'A', 26; i <= 'Z'; i, n = i+1, n+1 {
		alphabet = append(alphabet, byte(i))
		points[byte(i)] = n + 1
		appeard[byte(i)] = 0
	}

	for i, n := 'a', 0; i <= 'z'; i, n = i+1, n+1 {
		alphabet = append(alphabet, byte(i))
		points[byte(i)] = n + 1
		appeard[byte(i)] = 0
	}

	res := 0
	letters := []byte{}

	for _, line := range bytes.Split(data, []byte("\n")) {
		left, right := line[:len(line)/2], line[len(line)/2:]

		for _, v := range line {
			appeard[v]++
		}

		for k, value := range appeard {
			if value >= 2 {
				letters = append(letters, k)
			}
			appeard[k] = 0
		}

		for _, v := range letters {
			if bytes.ContainsRune(left, rune(v)) && bytes.ContainsRune(right, rune(v)) {
				res += points[v]
			}
		}

		// clear stuff out
		letters = []byte{}
	}

	fmt.Println(res)
}

// 7763
