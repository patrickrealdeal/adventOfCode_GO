package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(data)

	dim := len(strings.Split(lines, "\n"))
	numbers := make([][]int, dim)

	// populate matrix
	for i, line := range strings.Split(lines, "\n") {
		for j := 0; j < dim; j++ {
			numbers[i] = append(numbers[i], int(line[j]))
		}
	}

	res := 0

	for i := range numbers {
		for j := range numbers {
			counter := 0
			score := 1
			current := numbers[i][j]

			// Rows
			// right
			for r := j + 1; r < dim; r++ {
				counter++
				if numbers[i][r] >= current {
					break
				}
			}
			score *= counter
			counter = 0

			// left
			for l := j - 1; l >= 0; l-- {
				counter++
				if numbers[i][l] >= current {
					break
				}
			}
			score *= counter
			counter = 0

			// Cols
			// down
			for d := i + 1; d < dim; d++ {
				counter++
				if numbers[d][j] >= current {
					break
				}
			}
			score *= counter
			counter = 0

			// up
			for u := i - 1; u >= 0; u-- {
				counter++
				if numbers[u][j] >= current {
					break
				}
			}
			score *= counter

			if score > res {
				res = score
			}
		}
	}

	fmt.Println(res)
}

// 321975
