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

	uppercase := []byte{}
	lowercase := []byte{}
	points := make(map[byte]int)
	appeard := make(map[byte]int)

	for i, n := 'A', 26; i <= 'Z'; i, n = i+1, n+1 {
		uppercase = append(uppercase, byte(i))
		points[byte(i)] = n + 1
		appeard[byte(i)] = 0
	}

	for i, n := 'a', 0; i <= 'z'; i, n = i+1, n+1 {
		lowercase = append(lowercase, byte(i))
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
					if _, ok := appeard[c]; ok {
						appeard[c]++
					}
				}
			}

			for _, lower := range lowercase {
				if value, ok := appeard[lower]; ok {
					if value >= 3 {
						letters = append(letters, lower)
					}
				}
				appeard[lower] = 0
			}

			for _, upper := range uppercase {
				if value, ok := appeard[upper]; ok {
					if value >= 3 {
						letters = append(letters, upper)
					}
				}
				appeard[upper] = 0
			}

			for _, v := range letters {
				if bytes.Contains(block[0], []byte(string(v))) &&
					bytes.Contains(block[1], []byte(string(v))) &&
					bytes.Contains(block[2], []byte(string(v))) {
					res += points[v]
				}
			}
			block = [][]byte{}
			letters = []byte{}
		}
	}
	fmt.Println(res)
}

// 2569
