package main

import (
	"math/rand"
	"sort"
)

type Config struct {
	p int // number of partition
	max_b int // maximum number of data blocks in each partition
	min_b int // minimum number of data blocks in each partition
	max_e int // maximum number of elements in each block
	min_e int // minimum number of elements in each block
}

type ElType interface {
	LessThan(*ElType) bool
	IsEqualTo(*ElType) bool
	Init()
	ToStr() string
}

type Element struct {
	data int
}

func (lhs *Element) LessThan(rhs *ElType) bool {
	//var c = (*rhs).(Element)
	var d = (*rhs).(*Element)
	return lhs.data < d.data
}

func (lhs *Element) IsEqualTo(rhs *ElType) bool {
	var d = (*rhs).(*Element)
	return lhs.data == d.data
	//return lhs.data == (*rhs).(Element).data
}

func (lhs *Element) Init() {
	lhs.data = rand.Int()
}

func (lhs *Element) ToStr() string {
	return string(lhs.data)
}

func NewElement() ElType {
	var e = &Element{}
	e.Init()
	return e
}

type DataBlock []ElType

func (a DataBlock) Len() int {
	return len(a)
}

func (a DataBlock) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a DataBlock) Less(i, j int) bool {
	return a[i].LessThan(&a[j])
}

type Partition []DataBlock

func Generate(c Config, new func() ElType) *[]Partition {
	var parts []Partition
	parts = make([]Partition, c.p)

	for i := 0; i < c.p; i++ {
		var data DataBlock
		b := c.min_b + rand.Intn(c.max_b - c.min_b + 1)
		tot := 0
		sizes := make([]int, b + 1)
		for j := 0; j < b; j++ {
			e := c.min_e + rand.Intn(c.max_e-c.min_e+1)
			sizes[j + 1] = sizes[j] + e
			tot += e
		}
		sizes[b] = tot
		data = make([]ElType, tot)
		for j := 0; j < tot; j++ {
			data[j] = new()
		}
		sort.Sort(data)
		part := Partition{}
		for j := 0; j < b; j++ {
			var block = DataBlock{}
			block = data[sizes[j]:sizes[j + 1]]
			part = append(part, block)
		}
		parts = append(parts, part)
	}
	return &parts
}

func (p *Partition) Sort(method string) {
	switch method {
	case "bruteforce":

	}
}