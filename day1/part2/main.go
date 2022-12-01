// Part 2 of adv-2022 day1
package main

import (
	"fmt"
	"log"
	"os"
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
		for _, v := range d.cals {
			sum += v.amount
		}
		d.total = sum
	}

	max := 0
	maxCalorie := &dwarf{}
	top3 := []*dwarf{}
	sum := 0

	for len(top3) < 3 {
		for i, e := range dwarves {
			if i == 0 || e.total > max {
				max = e.total
				maxCalorie = e
			}
		}

		top3 = append(top3, maxCalorie)

		for i, c := range dwarves {
			if c.total == maxCalorie.total {
				dwarves = append(dwarves[:i], dwarves[i+1:]...)
			}
		}
		maxCalorie = &dwarf{}
	}

	for _, v := range top3 {
		sum += v.total
	}

	fmt.Println(sum)
}
