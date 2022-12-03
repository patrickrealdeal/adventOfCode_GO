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
	block := [][]byte{}

	for i, line := range bytes.Split(data, []byte("\n")) {
		block = append(block, line)

		if i%3 == 2 {
			for _, elem := range block {
				for _, c := range elem {
					appeard[c]++
				}
			}

			for k, value := range appeard {
				if value >= 3 {
					letters = append(letters, k)
				}
				appeard[k] = 0
			}

			for _, letter := range letters {
				if bytes.ContainsRune(block[0], rune(letter)) &&
					bytes.ContainsRune(block[1], rune(letter)) &&
					bytes.ContainsRune(block[2], rune(letter)) {
					res += points[letter]
				}
			}
			block = [][]byte{}
			letters = []byte{}
		}
	}
	fmt.Println(res)
}

// 2569
