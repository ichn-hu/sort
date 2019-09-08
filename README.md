# sort
Homework from PingCAP recruitment

# Description

Let's call a sorted array of numbers (or elements from a set that have partial order relationship defined) a data block.
Two data blocks are said to be non-overlapping if the ranges defined by the numbers in the data blocks are disjoint.
Sorted non-overlapping data blocks constitute a partition. Given several partitions, find a way to sort all the numbers 
appeared. You are encouraged to optimize your solution as to your best extent.

# Possible solution

## Bruteforce

Just concatenate all data blocks into a large array, then use any sorting algorithm to sort it.

## Restate as a k way merge problem

Since data blocks within a partition is ordered, each partition can be seen as a sorted list, and then the problem 
becomes a well know problem of sorting k ordered list.

## Heap based k way merge sort

This is a classical problem. Given k sorted lists, figure out an ordered list of all the elements in the k sorted lists.
This can be down using a heap which maintains pointers to the first unmerged element in each list, i.e. the minimal 
unmerged element in each list, and the sorted property of each list ensures that the minimal element in the heap is the 
global minimal unmerged element. Whenever we take the minimal element from the top of the heap, we increase the value of
that node to the next element in the corresponding list, and sink the node down in order to main the heap property. The
complexity of the algorithm is O(n log k), and it is optimal for k way merging, if better complexity exists, we could 
use it against sorting k element and obtain less than O(k log k) sorting algorithm, which contradicts that the optimal 
sorting complexity based on comparison is bounded by O(n log n).

## Tournament tree based merge sort

Tournament tree is a tree structure that maintains a complete binary tree as a competing hierarchy among k elements. If we maintain the winner (in our case, with smaller value) of 
each competition up until only one element remained, it will be the minimal element, and each time we remove the minimal element, we could 
update from leaf node up until root node to maintain the tree structure. We put each list a the leaf node, and maintain a pointer similarly.

The algorithm runs in O(n log k) as well, since each update takes O(tree depth) time, and since it is a complete binary tree, tree depth is O(log k).


# Benchmark

```text
(base)  ✘ ichn@ichn-arch-pc  ~/Projects/sort/src/sort   master ●✚  go test -bench=.
goos: linux
goarch: amd64
BenchmarkSortKWayMergeViaBinaryHeapRandomP10B1000E1000-8                               1        3990439316 ns/op
BenchmarkSortKWayMergeViaBinaryHeapRandomP1000B100E100-8                               1        4768437472 ns/op
BenchmarkSortKWayMergeViaBinaryHeapRandomP100000B10E10-8                               1        17305088317 ns/op
BenchmarkSortKWayMergeViaBinaryHeapRandomP10000000B1E1-8                               1        46030324661 ns/op
BenchmarkSortKWayMergeViaTournamentTreeRandomP10B1000E1000-8                           1        3047975517 ns/op
BenchmarkSortKWayMergeViaTournamentTreeRandomP1000B100E100-8                           1        4644232134 ns/op
BenchmarkSortKWayMergeViaTournamentTreeRandomP100000B10E10-8                           1        16841555829 ns/op
BenchmarkSortKWayMergeViaTournamentTreeRandomP10000000B1E1-8                           1        36817441663 ns/op
BenchmarkSortKWayMergeViaBinaryHeapNonOverlappingP10B1000E1000-8                       1        1301352322 ns/op
BenchmarkSortKWayMergeViaTournamentTreeNonOverlappingP10B1000E1000-8                   1        1164388819 ns/op
BenchmarkSortKWayMergeViaBinaryHeapNonOverlappingP100000B10E10-8                       1        2816296290 ns/op
BenchmarkSortKWayMergeViaTournamentTreeNonOverlappingP100000B10E10-8                   1        4599237971 ns/op
BenchmarkSortKWayMergeViaBinaryHeapNonOverlappingP1000000B10E10-8                      1        49949588686 ns/op
BenchmarkSortKWayMergeViaTournamentTreeNonOverlappingP1000000B10E10-8                  1        100163527913 ns/op
PASS
ok      _/home/ichn/Projects/sort/src/sort      352.203s
```

In the benchmark, P refers to the number of partitions, B means how many blocks will there be in each partition, and E 
stands for number of elements in each data block.

A exhaust benchmarking resulted in several observations. 

* Tournament tree based method outperforms heap if data is generated randomly
* If data blocks across partition are not overlapped, then heap performs better

# Comparision

Using tournament tree to maintain the minimal element could result in less memory access, since for each update, it only requires 1 comparision with a sibling node,
while to maintain the heap, floating one node up or sink it down for one level requires 2 comparision, therefore overral using tournament tree could result in less running time.

However, if the data blocks across different partition also don't overlap, then using heap will be more favorable, because for the update within a same data block,
the minimal node won't sink in the heap, therefore each update only result in 2 comparisions. However for tournament tree it still need a bottom up update, which is O (log k) no matter how
the data is structured. 

# Conclusion

If data across partitions also don't overlap, we could just sort the first element of each data block, and would result in a 
algorithm the runs in O(P log B + n). However this problem turns out to be a totally different problem, and won't scale to situation where the non-overlapping guarantee does not exists.

For simplicity and because of time limitation for the homework, I have not yet developed a hybrid algorithm that could solve the problem with or without the condition,
but I still believe the heap solution is quite acceptable, for its elegance and scalability, and the tournament tree solution can be an alternative when there is no any prior knowledge in the distribution of data cross partitions.