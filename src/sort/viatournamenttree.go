package main

type TreeNode struct {
	d DataBlock
	idx int
	from int
}

type Tree []TreeNode

func (t Tree) competeLeaf(l, r int) int {
	if t[l].d[t[l].idx].LessThan(&t[r].d[t[r].idx]) {
		return l
	}
	return r
}

func (t Tree) notNil(i int) bool {
	return t[i].from != 0
}

func (t Tree) extractMin() ElType {
	leaf := &t[t[1].from]
	ret := leaf.d[leaf.idx]
	leaf.idx++
	if leaf.idx == leaf.d.Len() {
		leaf.from = 0
	}
	for i := t[1].from / 2; i != 0; i /= 2 {
		t.updateNode(i)
	}
	return ret
}

func (t Tree) updateNode(x int) {
	left := x * 2
	right := left + 1
	if t.notNil(left) || t.notNil(right) {
		if t.notNil(left) {
			if t.notNil(right) {
				t[x].from = t.competeLeaf(t[left].from, t[right].from)
			} else {
				t[x].from = t[left].from
			}
		} else {
			t[x].from = t[right].from
		}
	} else {
		t[x].from = 0
	}
}

func SortKWayMergeViaTournamentTree(parts *[]Partition) DataBlock {
	numParts := len(*parts)
	level, capacity := 0, 1
	for ; capacity < numParts;  {
		level = level + 1
		capacity = capacity * 2
	}
	var tree Tree = make([]TreeNode, capacity* 2)
	tot := 0
	for i := 0; i < numParts; i++ {
		leaf := &tree[i + capacity]
		leaf.d = (*parts)[i].Squash()
		leaf.idx = 0
		if leaf.d.Len() != 0 {
			leaf.from = i + capacity
			tot += leaf.d.Len()
		}
	}
	for i, j := capacity/ 2, capacity; i != 0; i, j = i / 2, j / 2 {
		for l := i; l < j; l++ {
			tree.updateNode(l)
		}
	}
	res := DataBlock{}
	res = make([]ElType, tot)
	for i := 0; i < tot; i++ {
		res[i] = tree.extractMin()
	}
	return res
}
