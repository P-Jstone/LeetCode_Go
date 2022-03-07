package GreedyAlgorithm

import "sort"

func findMinArrowShots(points [][]int) int {
	res := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	end := points[0][1]
	for _, p := range points {
		if p[0] > end {
			res++
			end = p[1]
		}
	}
	return res
}
