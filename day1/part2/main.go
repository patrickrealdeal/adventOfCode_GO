// Part 2 of adv-2022 day1
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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
	buf, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(buf)
	dwarves := []*dwarf{}
	cals := []*calories{}

	lines := strings.Split(s, "\n\n")
	for _, line := range lines {
		elems := strings.Split(strings.TrimSpace(line), "\n")

		for _, e := range elems {
			n, err := strconv.Atoi(e)
			if err != nil {
				log.Fatal(err)
			}

			c := newCalorie(n)
			cals = append(cals, c)
		}

		d := newDwarf()
		d.cals = cals
		dwarves = append(dwarves, d)

		cals = []*calories{}
	}

	for _, d := range dwarves {
		sum := 0
		for _, c := range d.cals {
			sum += c.amount
		}
		d.total = sum
	}

	sum := 0
	total := []int{}
	for _, d := range dwarves {
		total = append(total, d.total)
	}

	sort.Ints(total)
	res := total[len(total)-3:]
	for _, v := range res {
		sum += v
	}

	fmt.Println(sum)
}
