package Pointer

import "math"

func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		if l*l+r*r < c {
			l++
		} else if l*l+r*r > c {
			r--
		} else {
			return true
		}
	}
	return false
}
