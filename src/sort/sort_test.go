package main

import (
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
	if !valid(GenerateRandom(c, NewElement)) {
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
		if !valid(GenerateRandom(c, NewElement)) {
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
	_ = SortBruteforce(GenerateRandom(c, NewElement))
}

func TestSortKWayMergeViaBinaryHeap(t *testing.T) {
	round := 1
	for r := 0; r < round; r++ {
		c := Config{
			p:     rand.Intn(100),
			max_b: rand.Intn(1000),
			min_b: 0,
			max_e: rand.Intn(1000),
			min_e: 0,
		}
		d := GenerateRandom(c, NewElement)
		gold := SortBruteforce(d)
		result := SortKWayMergeViaBinaryHeap(d)
		if !gold.Equals(result) {
			t.Fatal("GG", gold, result)
		}
	}
}

func TestSortKWayMergeViaLoserTree(t *testing.T) {
	round := 1
	for r := 0; r < round; r++ {
		c := Config{
			p:     rand.Intn(100),
			max_b: rand.Intn(1000),
			min_b: 0,
			max_e: rand.Intn(1000),
			min_e: 0,
		}
		d := GenerateRandom(c, NewElement)
		gold := SortBruteforce(d)
		result := SortKWayMergeViaLoserTree(d)
		if !gold.Equals(result) {
			t.Fatal("GG", gold, result)
		}
	}
}

func BenchmarkSortKWayMergeViaBinaryHeapRandomP10B1000E1000(b *testing.B) {
	c := Config{
		p:     10,
		max_b: 1000,
		min_b: 1000,
		max_e: 1000,
		min_e: 1000,
	}
	d := GenerateRandom(c, NewElement)
	_ = SortKWayMergeViaBinaryHeap(d)
}

func BenchmarkSortKWayMergeViaBinaryHeapRandomP1000B100E100(b *testing.B) {
	c := Config{
		p:     1000,
		max_b: 100,
		min_b: 100,
		max_e: 100,
		min_e: 100,
	}
	d := GenerateRandom(c, NewElement)
	_ = SortKWayMergeViaBinaryHeap(d)
}

func BenchmarkSortKWayMergeViaBinaryHeapRandomP100000B10E10(b *testing.B) {
	c := Config{
		p:     100000,
		max_b: 10,
		min_b: 10,
		max_e: 10,
		min_e: 10,
	}
	d := GenerateRandom(c, NewElement)
	_ = SortKWayMergeViaBinaryHeap(d)
}

func BenchmarkSortKWayMergeViaBinaryHeapRandomP10000000B1E1(b *testing.B) {
	c := Config{
		p:     10000000,
		max_b: 1,
		min_b: 1,
		max_e: 1,
		min_e: 1,
	}
	d := GenerateRandom(c, NewElement)
	_ = SortKWayMergeViaBinaryHeap(d)
}

func BenchmarkSortKWayMergeViaLoserTreeRandomP10B1000E1000(b *testing.B) {
	c := Config{
		p:     10,
		max_b: 1000,
		min_b: 1000,
		max_e: 1000,
		min_e: 1000,
	}
	d := GenerateRandom(c, NewElement)
	_ = SortKWayMergeViaLoserTree(d)
}

func BenchmarkSortKWayMergeViaLoserTreeRandomP1000B100E100(b *testing.B) {
	c := Config{
		p:     1000,
		max_b: 100,
		min_b: 100,
		max_e: 100,
		min_e: 100,
	}
	d := GenerateRandom(c, NewElement)
	_ = SortKWayMergeViaLoserTree(d)
}

func BenchmarkSortKWayMergeViaLoserTreeRandomP100000B10E10(b *testing.B) {
	c := Config{
		p:     100000,
		max_b: 10,
		min_b: 10,
		max_e: 10,
		min_e: 10,
	}
	d := GenerateRandom(c, NewElement)
	_ = SortKWayMergeViaLoserTree(d)
}

func BenchmarkSortKWayMergeViaLoserTreeRandomP10000000B1E1(b *testing.B) {
	c := Config{
		p:     10000000,
		max_b: 1,
		min_b: 1,
		max_e: 1,
		min_e: 1,
	}
	d := GenerateRandom(c, NewElement)
	_ = SortKWayMergeViaLoserTree(d)
}

func BenchmarkSortKWayMergeViaBinaryHeapNonOverlappingP10B1000E1000(b *testing.B) {
	c := Config{
		p:     10,
		max_b: 1000,
		min_b: 1000,
		max_e: 1000,
		min_e: 1000,
	}
	d := GenerateNonOverlapping(c, NewElement)
	SortKWayMergeViaBinaryHeap(d)
}

func BenchmarkSortKWayMergeViaLoserTreeNonOverlappingP10B1000E1000(b *testing.B) {
	c := Config{
		p:     10,
		max_b: 1000,
		min_b: 1000,
		max_e: 1000,
		min_e: 1000,
	}
	d := GenerateNonOverlapping(c, NewElement)
	SortKWayMergeViaLoserTree(d)
}

func BenchmarkSortKWayMergeViaBinaryHeapNonOverlappingP100000B10E10(b *testing.B) {
	c := Config{
		p:     100000,
		max_b: 10,
		min_b: 10,
		max_e: 10,
		min_e: 10,
	}
	d := GenerateNonOverlapping(c, NewElement)
	SortKWayMergeViaBinaryHeap(d)
}

func BenchmarkSortKWayMergeViaLoserTreeNonOverlappingP100000B10E10(b *testing.B) {
	c := Config{
		p:     100000,
		max_b: 10,
		min_b: 10,
		max_e: 10,
		min_e: 10,
	}
	d := GenerateNonOverlapping(c, NewElement)
	SortKWayMergeViaLoserTree(d)
}

/*

BenchmarkSortKWayMergeViaBinaryHeapRandomP10B1000E1000-8               1        13447202900 ns/op
BenchmarkSortKWayMergeViaBinaryHeapRandomP1000B100E100-8               1        11210108595 ns/op
BenchmarkSortKWayMergeViaBinaryHeapRandomP100000B10E10-8               1        37232284421 ns/op
BenchmarkSortKWayMergeViaBinaryHeapRandomP10000000B1E1-8               1        77178133630 ns/op
BenchmarkSortKWayMergeViaLoserTreeRandomP10B1000E1000-8                1        12377093275 ns/op
BenchmarkSortKWayMergeViaLoserTreeRandomP1000B100E100-8                1        9653308611 ns/op
BenchmarkSortKWayMergeViaLoserTreeRandomP100000B10E10-8                1        35112465734 ns/op
BenchmarkSortKWayMergeViaLoserTreeRandomP10000000B1E1-8                1        61994628846 ns/op
13447202900
17719069552

BenchmarkSortKWayMergeViaBinaryHeapNonOverlappingP10B1000E1000-8               1        17719069552 ns/op
BenchmarkSortKWayMergeViaLoserTreeNonOverlappingP10B1000E1000-8                1        17658385611 ns/op

 */


