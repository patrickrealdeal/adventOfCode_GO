// Part 2 of adv-2022 day1
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type dwarf struct {
	cals  []*calories
	total int
}

type calories struct {
	amount int
}

func (d *dwarf) Calories() []*calories {
	return d.cals
}

func newDwarf() *dwarf {
	return &dwarf{cals: []*calories{}}
}

func newCalorie(amount int) *calories {
	return &calories{
		amount: amount,
	}
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dwarves := []*dwarf{}
	cals := []*calories{}

	for _, line := range bytes.Split(data, []byte("\n")) {
		if string(line) == "" {
			d := newDwarf()
			d.cals = cals
			dwarves = append(dwarves, d)

			cals = []*calories{}
			continue
		}

		n, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatal(err)
		}

		c := newCalorie(n)
		cals = append(cals, c)
	}

	total := []int{}
	for _, d := range dwarves {
		var sum int
		for _, c := range d.cals {
			sum += c.amount
		}
		d.total = sum
		total = append(total, d.total)
	}

	var sum int
	sort.Ints(total)
	res := total[len(total)-3:]
	for _, v := range res {
		sum += v
	}

	fmt.Println(sum)
}
