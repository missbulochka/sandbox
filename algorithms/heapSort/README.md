# Heap sort

## Description

Heap sort is an efficient sort algorithm.
Efficiency in the best, average and worst case is **O(n log n)**.

## Algorithm

The idea of ​​heapsort is to use a data structure such as a binary heap,
which is a binary tree subject to the following restrictions:
1.  The value in any node is not less than in any of its children
2.  The binary tree is solid (each node has 0 or 2 children)

To represent a tree through an array, you can use the following idea:
the first position is the root of the binary tree, then for each child it
will be true that its index is defined as 2n + 1 for the left child and
2n + 2 for the right child:

![Heap-as-array.svg](Heap-as-array.svg)

[Visualisation](https://www.toptal.com/developers/sorting-algorithms/heap-sort)

The heap sort consists of the following steps:
1.  The elements of the input sequence must be rearranged so that they satisfy the conditions for a binary heap
2.  Exchange the first element with the last, remove the element from data consideration and perform an operation
to restore the binary heap, because after the exchange, the resulting tree may not meet constraint 1
