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


```

# Profiling
