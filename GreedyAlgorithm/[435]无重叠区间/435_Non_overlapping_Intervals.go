package GreedyAlgorithm

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	preserve, prev := 1, intervals[0][1]
	for _, v := range intervals {
		if v[0] >= prev {
			prev = v[1]
			preserve++
		}
	}
	return n - preserve
}
