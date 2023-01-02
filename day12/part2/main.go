package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	c [2]int
	v int
}

var dirs = [][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := string(data)

	var seen = map[[2]int]bool{}
	_, start, board := LoadBoard(lines)

	// Queue for positions we want to evaluate
	queue := make(chan Coord, len(board)*len(board[0]))

	// Starting position
	queue <- Coord{start, 0}

	// While we have positions to evaluate
	for len(queue) > 0 {
		cur := <-queue
		// Check if we reached the end
		if board[cur.c[0]][cur.c[1]] == 0 {
			fmt.Println(cur.v)
			return
		}

		// Check if we already evaluated this earlier
		if seen[cur.c] {
			continue
		}

		seen[cur.c] = true
		// Check all directions
		for _, dir := range dirs {
			nx, ny := cur.c[0]+dir[0], cur.c[1]+dir[1]
			// Make sure we're inside the board
			if nx < 0 || ny < 0 || nx >= len(board) || ny >= len(board[0]) {
				continue
			}
			// Make sure we are making legal moves
			if !canGo(board[nx][ny], board[cur.c[0]][cur.c[1]]) {
				continue
			}
			// Add to evaluation list
			queue <- Coord{[2]int{nx, ny}, cur.v + 1}
		}
	}
}

func LoadBoard(lines string) (start, end [2]int, board [][]int) {
	// Loading map and finding start and end
	for i, line := range strings.Split(lines, "\n") {
		var curLine = make([]int, len(line))

		for j := range line {
			if line[j] == 'S' { // Start is 0 elevation
				start = [2]int{i, j}
				curLine[j] = 0 // elevation
			} else if line[j] == 'E' { // End is 25 elevation
				end = [2]int{i, j}
				curLine[j] = 25 // elevation
			} else {
				curLine[j] = int(line[j] - 'a')
			}
		}
		board = append(board, curLine)
	}
	return start, end, board
}

// Can go from a to b
func canGo(a, b int) bool {
	if a > b {
		return true
	}
	return b-a <= 1
}
