package main

import "container/heap"

type HeapNode struct {
	val ElType
	idx int
	d DataBlock
}

// nexVal turn val into next value in the data block, return false if the current value is the last one
func (n *HeapNode) nextVal() bool {
	n.idx++
	if n.idx == n.d.Len() {
		return false
	}
	n.val = n.d[n.idx]
	return true
}

type Heap []HeapNode

func (h *Heap) Less(i, j int) bool {
	return (*h)[i].val.LessThan(&(*h)[j].val)
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(HeapNode))
}

func (h *Heap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *Heap) Empty() bool {
	return h.Len() == 0
}

func (h *Heap) Min() *HeapNode {
	return &(*h)[0]
}

func SortKWayMergeViaBinaryHeap(parts *[]Partition) DataBlock {
	h := new(Heap)
	ret := DataBlock{}
	for _, p := range *parts {
		d := p.Squash()
		if len(d) != 0 {
			*h = append(*h, HeapNode{
				val: d[0],
				idx: 0,
				d:   d,
			})
		}
	}
	heap.Init(h)
	for !h.Empty() {
		minNode := h.Min()
		ret = append(ret, minNode.val)
		if minNode.nextVal() {
			heap.Fix(h, 0)
		} else {
			heap.Remove(h, 0)
		}
	}
	return ret
}