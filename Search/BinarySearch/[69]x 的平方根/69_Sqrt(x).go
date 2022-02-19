package Search

func mySqrt(x int) int {
	var y int
	y = x
	for y*y > x {
		y = (y + x/y) / 2
	}
	return y
}
