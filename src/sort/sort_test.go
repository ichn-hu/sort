package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGenerate(t *testing.T) {
	valid := func(parts *[]Partition) bool {
		for _, part := range *parts {
			var pivot *Element = nil
			for _, block := range part {
				for _, el := range block {
					if pivot == nil {
						pivot = el.(*Element)
					} else {
						if !pivot.LessThan(&el) {
							return false
						}
					}
				}
			}
		}
		return true
	}

	c := Config{
		p:     0,
		max_b: 0,
		min_b: 0,
		max_e: 0,
		min_e: 0,
	}
	if !valid(Generate(c, NewElement)) {
		t.Fatal("Generated result not valid")
	}

	round := 100
	for r := 0; r < round; r++ {
		c := Config{
			p:     rand.Intn(100),
			max_b: rand.Intn(100),
			min_b: 0,
			max_e: rand.Intn(100),
			min_e: 0,
		}
		if !valid(Generate(c, NewElement)) {
			t.Fatal("Generated result not valid")
		}
	}
}

func TestSortBruteforce(t *testing.T) {
	c := Config{
		p:     10,
		max_b: 100,
		min_b: 0,
		max_e: 100,
		min_e: 0,
	}
	d := SortBruteforce(Generate(c, NewElement))
	fmt.Print(d)
}
