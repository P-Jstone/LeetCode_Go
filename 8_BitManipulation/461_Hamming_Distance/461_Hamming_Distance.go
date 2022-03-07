package BM

func hammingDistance(x int, y int) int {
	var res int
	for s := x ^ y; s > 0; s &= s - 1 {
		res++
	}
	return res
}
