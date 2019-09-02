package main

import (
	"fmt"
	"sort"
)

func SortBruteforce(parts *[]Partition) DataBlock {
	var data DataBlock
	for _, part := range *parts {
		for _, block := range part {
			for _, el := range block {
				data = append(data, el)
			}
		}
	}
	sort.Sort(data)
	return data
}




func main() {
	var c = Config{
		p:     3,
		max_b: 3,
		min_b: 2,
		max_e: 3,
		min_e: 2,
	}
	var p = Generate(c, NewElement)
	fmt.Print(p)
}
