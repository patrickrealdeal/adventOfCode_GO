package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(data)

	dim := len(strings.Split(lines, "\n"))
	visibles := make([][]int, dim)
	numbers := make([][]int, dim)
	res := 0

	// populate matrix
	for i, line := range strings.Split(lines, "\n") {
		for j := 0; j < dim; j++ {
			numbers[i] = append(numbers[i], int(line[j]))
			visibles[i] = append(visibles[i], 0)
		}
	}

	// rows
	for i := 0; i < dim; i++ {
		max := -1

		// right
		for r := 0; r < dim; r++ {
			if numbers[i][r] > max {
				visibles[i][r] = 1
				max = numbers[i][r]
				continue
			}
		}

		max = -1

		//left
		for l := dim - 1; l > 0; l-- {
			if numbers[i][l] > max {
				visibles[i][l] = 1
				max = numbers[i][l]
				continue
			}
		}
	}

	// cols
	for j := 0; j < dim; j++ {
		max := -1

		// down
		for d := 0; d < dim; d++ {
			if numbers[d][j] > max {
				visibles[d][j] = 1
				max = numbers[d][j]
				continue
			}
		}

		max = -1

		// up
		for u := dim - 1; u > 0; u-- {
			if numbers[u][j] > max {
				visibles[u][j] = 1
				max = numbers[u][j]
				continue
			}
		}
	}

	for i := 0; i < len(visibles); i++ {
		for j := 0; j < len(visibles); j++ {
			if visibles[i][j] == 1 {
				res++
			}
		}
	}

	fmt.Println(res)
}

// 1717
