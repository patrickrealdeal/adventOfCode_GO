package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	N         uint64
	Items     []uint64
	Operation func(uint64) uint64
	Test      uint64
	Pos       uint64
	Neg       uint64
	Inspected uint64
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := string(data)

	var total uint64
	var op func(uint64) uint64
	monkeys := []*monkey{}
	line := strings.Split(lines, "\n\n")

	for _, v := range line {
		line := strings.Split(v, "\n")
		n, _ := strconv.ParseUint(string(line[0][7]), 10, 64)
		ns := strings.Split(line[1][18:], ", ")
		items := []uint64{}

		for _, v := range ns {
			item, _ := strconv.ParseUint(v, 10, 64)
			items = append(items, item)
		}

		ops := strings.Split(line[2][23:], " ")

		if strings.Contains(ops[1], "old") {
			op = func(u uint64) uint64 {
				if ops[0] == "*" {
					return u * u
				}
				return u + u
			}
		} else {
			val, _ := strconv.ParseUint(ops[1], 10, 64)
			op = func(a uint64) uint64 {
				if ops[0] == "*" {
					return a * val
				}
				return a + val
			}
		}

		test, _ := strconv.ParseUint(string(line[3][21:]), 10, 64)
		pos, _ := strconv.ParseUint(string(line[4][len(line[4])-1]), 10, 64)
		neg, _ := strconv.ParseUint(string(line[5][len(line[5])-1]), 10, 64)

		monkey := &monkey{
			N:         n,
			Items:     items,
			Operation: op,
			Pos:       pos,
			Neg:       neg,
			Test:      test,
		}

		monkeys = append(monkeys, monkey)
	}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for _, w := range monkey.Items {
				w = monkey.Operation(w) / 3

				if w%monkey.Test == 0 {
					monkeys[monkey.Pos].Items = append(monkeys[monkey.Pos].Items, w)
					monkey.Items = monkey.Items[1:]
				} else {
					monkeys[monkey.Neg].Items = append(monkeys[monkey.Neg].Items, w)
					monkey.Items = monkey.Items[1:]
				}
				monkey.Inspected++
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})

	total = monkeys[0].Inspected * monkeys[1].Inspected

	fmt.Println(total)
}

// 120384
