package main

type LoserTreeNode struct {
	loser *Data
	winner *Data
}

type Data struct {
	id int
	d DataBlock
	idx int
	size int
	val ElType
}

type LoserTree struct {
	tree []LoserTreeNode
	data []Data
	largest ElType
}

func (t *LoserTree) updateNode(x int) {
	fa := x / 2
	//log.Println("Update ", x, " fa ", fa)
	fn := &t.tree[fa]
	xx := &t.tree[x]
	if xx.winner == nil {
		//log.Println(x, t.tree)
		return
	}
	xval := xx.winner.val
	if fn.loser != nil {
		if fn.loser.val.LessThan(&xval) {
			fn.winner = fn.loser
			fn.loser = xx.winner
		} else {
			fn.winner = xx.winner
		}
	} else {
		fn.loser = xx.winner
		fn.winner = xx.winner
	}
}

func (t *LoserTree) extractMin() ElType {
	winner := t.tree[0].winner
	res := winner.val
	winner.idx++
	if winner.idx >= winner.size {
		winner.val = t.largest
	} else {
		winner.val = winner.d[winner.idx]
	}
	for i := winner.id; 0 < i; i/=2 {
		t.updateNode(i)
	}
	return res
}

func SortKWayMergeViaLoserTree(parts *[]Partition) DataBlock {
	data := make([]Data, 0)
	tot := 0
	var largest ElType = nil
	for i, j := 0, len(*parts); i < j; i++ {
		d := (*parts)[i].Squash()
		size := d.Len()
		if size > 0 {
			if largest == nil {
				largest = d[0]
			}
			for i := 0; i < size; i++ {
				if largest.LessThan(&d[i]) {
					largest = d[i]
				}
			}
			data = append(data, Data{0, d, 0, size, d[0]})
			tot += size
		}
	}
	numParts := len(data)
	level, capacity := 0, 1
	for ; capacity < numParts;  {
		level = level + 1
		capacity = capacity * 2
	}
	t := LoserTree{
		tree: make([]LoserTreeNode, capacity * 2),
		data: data,
		largest: largest,
	}
	for i := 0; i < numParts; i++ {
		id := i + capacity
		t.data[i].id = id
		t.tree[id] = LoserTreeNode{
			loser:  nil,
			winner: &t.data[i],
		}
		//log.Println("Update ", id, " ", i, " ", numParts, " ", t.tree)
		t.updateNode(id)
	}
	for i, j := capacity/ 2, capacity; i != 0; i, j = i / 2, j / 2 {
		for l := i; l < j; l++ {
			t.updateNode(l)
		}
	}
	//log.Println("tree ", t.tree)

	res := DataBlock{}
	res = make([]ElType, tot)
	for i := 0; i < tot; i++ {
		//log.Println("working on ", i)
		res[i] = t.extractMin()
	}
	return res
}
