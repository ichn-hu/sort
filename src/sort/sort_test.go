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

func checkOrdered(d DataBlock, t *testing.T) {
	for i := 1; i < d.Len(); i++ {
		if !d[i - 1].LessThan(&d[i]) {
			t.Fatal("Result not sorted")
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
	d := SortBruteforce(GenerateRandom(c, NewElement))
	checkOrdered(d, t)
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

func TestSortKWayMergeViaTournamentTree(t *testing.T) {
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
		result := SortKWayMergeViaTournamentTree(d)
		if !gold.Equals(result) {
			t.Fatal("GG", gold, result)
		}
	}
}


func benchmarkHeap(parts *[]Partition, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SortKWayMergeViaBinaryHeap(parts)
	}
}

func benchmarkTournamentTree(parts *[]Partition, b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SortKWayMergeViaTournamentTree(parts)
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
	benchmarkHeap(d, b)
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
	benchmarkHeap(d, b)
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
	benchmarkHeap(d, b)
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
	benchmarkHeap(d, b)
}

func BenchmarkSortKWayMergeViaTournamentTreeRandomP10B1000E1000(b *testing.B) {
	c := Config{
		p:     10,
		max_b: 1000,
		min_b: 1000,
		max_e: 1000,
		min_e: 1000,
	}
	d := GenerateRandom(c, NewElement)
	benchmarkTournamentTree(d, b)
}

func BenchmarkSortKWayMergeViaTournamentTreeRandomP1000B100E100(b *testing.B) {
	c := Config{
		p:     1000,
		max_b: 100,
		min_b: 100,
		max_e: 100,
		min_e: 100,
	}
	d := GenerateRandom(c, NewElement)
	benchmarkTournamentTree(d, b)
}

func BenchmarkSortKWayMergeViaTournamentTreeRandomP100000B10E10(b *testing.B) {
	c := Config{
		p:     100000,
		max_b: 10,
		min_b: 10,
		max_e: 10,
		min_e: 10,
	}
	d := GenerateRandom(c, NewElement)
	benchmarkTournamentTree(d, b)
}

func BenchmarkSortKWayMergeViaTournamentTreeRandomP10000000B1E1(b *testing.B) {
	c := Config{
		p:     10000000,
		max_b: 1,
		min_b: 1,
		max_e: 1,
		min_e: 1,
	}
	d := GenerateRandom(c, NewElement)
	benchmarkTournamentTree(d, b)
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
	benchmarkHeap(d, b)
}

func BenchmarkSortKWayMergeViaTournamentTreeNonOverlappingP10B1000E1000(b *testing.B) {
	c := Config{
		p:     10,
		max_b: 1000,
		min_b: 1000,
		max_e: 1000,
		min_e: 1000,
	}
	d := GenerateNonOverlapping(c, NewElement)
	benchmarkTournamentTree(d, b)
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
	benchmarkHeap(d, b)
}

func BenchmarkSortKWayMergeViaTournamentTreeNonOverlappingP100000B10E10(b *testing.B) {
	c := Config{
		p:     100000,
		max_b: 10,
		min_b: 10,
		max_e: 10,
		min_e: 10,
	}
	d := GenerateNonOverlapping(c, NewElement)
	benchmarkTournamentTree(d, b)
}

func BenchmarkSortKWayMergeViaBinaryHeapNonOverlappingP1000000B10E10(b *testing.B) {
	c := Config{
		p:     1000000,
		max_b: 10,
		min_b: 10,
		max_e: 10,
		min_e: 10,
	}
	d := GenerateNonOverlapping(c, NewElement)
	benchmarkHeap(d, b)
}

func BenchmarkSortKWayMergeViaTournamentTreeNonOverlappingP1000000B10E10(b *testing.B) {
	c := Config{
		p:     1000000,
		max_b: 10,
		min_b: 10,
		max_e: 10,
		min_e: 10,
	}
	d := GenerateNonOverlapping(c, NewElement)
	benchmarkTournamentTree(d, b)
}
