package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := string(data)

	chunk := 14
	res := processed(line, chunk)
	fmt.Println(res)
}

func processed(line string, chunk int) int {
	for i := 0; i < len(line)-chunk; i++ {
		check := line[i : i+chunk]
		uniques := make(map[rune]bool)

		for _, c := range check {
			uniques[c] = true
		}

		if len(uniques) == chunk {
			res := i + chunk
			return res
		}
	}

	return 0
}

// 2265
