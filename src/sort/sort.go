package main

import (
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

