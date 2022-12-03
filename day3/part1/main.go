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
	points := make(map[byte]int)
	appeard := make(map[byte]int)

	for i, n := 'A', 26; i <= 'Z'; i, n = i+1, n+1 {
		uppercase = append(uppercase, byte(i))
		points[byte(i)] = n + 1
		appeard[byte(i)] = 0
	}

	lowercase := []byte{}
	for i, n := 'a', 0; i <= 'z'; i, n = i+1, n+1 {
		lowercase = append(lowercase, byte(i))
		points[byte(i)] = n + 1
		appeard[byte(i)] = 0
	}
	res := 0
	letters := []byte{}

	for _, line := range bytes.Split(data, []byte("\n")) {
		left, right := line[:len(line)/2], line[len(line)/2:]

		for _, v := range left {
			if _, ok := appeard[v]; ok {
				appeard[v]++
			}
		}

		for _, v := range right {
			if _, ok := appeard[v]; ok {
				appeard[v]++
			}
		}

		for _, v := range lowercase {
			if value, ok := appeard[v]; ok {
				if value > 1 {
					letters = append(letters, v)
				}
			}
			appeard[v] = 0
		}

		for _, v := range uppercase {
			if value, ok := appeard[v]; ok {
				if value > 1 {
					letters = append(letters, v)
				}
			}
			appeard[v] = 0
		}

		for _, v := range letters {
			if bytes.Contains(left, []byte(string(v))) && bytes.Contains(right, []byte(string(v))) {
				res += points[v]
			}
		}

		// clear stuff out
		letters = []byte{}
	}

	fmt.Println(res)
}

// 7763
