// Part 1 of adv-2022 day1
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
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

	c := []int{}
	cals := []*calories{}

	lines := bytes.Split(buf, []byte("\n\n"))
	for _, line := range lines {
		trim := bytes.Trim(line, "\n")
		elems := bytes.Split(trim, []byte("\n"))
		for _, e := range elems {
			trimmed := bytes.TrimSuffix(e, []byte("\n"))
			n, err := strconv.Atoi(string(trimmed))
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
	for i, e := range cals {
		if i == 0 || e.amount > max {
			max = e.amount
		}
	}

	fmt.Println(max)
}

// 70116
