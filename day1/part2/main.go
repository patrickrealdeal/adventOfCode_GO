// Part 2 of adv-2022 day1
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type calories struct {
	amount int
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
	c := []int{}
	cals := []*calories{}

	lines := strings.Split(s, "\n\n")
	for _, s := range lines {
		trim := strings.Trim(s, "\n")
		elems := strings.Split(trim, "\n")
		for _, e := range elems {
			trimmed := strings.TrimSuffix(e, "\n")
			n, err := strconv.Atoi(trimmed)
			if err != nil {
				log.Fatal(err)
			}

			c = append(c, n)
		}

		sum := 0
		for _, i := range c {
			sum += i
		}
		c = []int{}
		cal := newCalorie(sum)
		cals = append(cals, cal)
	}

	max := 0
	maxCalorie := &calories{}
	top3 := []*calories{}
	sum := 0

	for len(top3) < 3 {
		for i, e := range cals {
			if i == 0 || e.amount > max {
				max = e.amount
				maxCalorie = cals[i]
			}
		}

		top3 = append(top3, maxCalorie)

		for i, c := range cals {
			if c.amount == maxCalorie.amount {
				cals = append(cals[:i], cals[i+1:]...)
			}
		}
		maxCalorie = &calories{}
	}

	for _, v := range top3 {
		sum += v.amount
	}

	fmt.Println(sum)

}
