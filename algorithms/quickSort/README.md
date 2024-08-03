# Quick sort

## Description

Quick sort is an efficient sort algorithm.
The best and average case efficiency is **O(n log n)** but the worst case **O(n^2)**
(if the pivot for each sequence will be its last element).

## Algorithm

[Visualisation](https://www.toptal.com/developers/sorting-algorithms/quick-sort)

The quick sort method works on the principle of "divide & conquer":
1.  Choose the pivot.
2.  Splitting: Rearrange the order of sequence elements so that all elements greater
than or equal to the pivot element are after the pivot element (the sequences can be empty).
After separation, the pivot will be in its place.
3.  Recursively repeat the previous two steps.
