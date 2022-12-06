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

	chunk := 14
	res := processed(data, chunk)
	fmt.Println(res)
}

func processed(data []byte, chunk int) int {
	for i := 0; i < len(data)-chunk; i++ {
		check := data[i : i+chunk]
		uniques := make(map[byte]struct{}, chunk)
		for _, c := range check {
			uniques[c] = struct{}{}
		}

		if len(uniques) == chunk {
			res := i + chunk
			return res
		}
	}

	return -1
}

// 2265
