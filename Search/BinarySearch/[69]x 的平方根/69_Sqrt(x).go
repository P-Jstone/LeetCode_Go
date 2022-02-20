package Search

// 二分查找
func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + ((r - l) >> 1)
		if mid*mid <= x {
			if (mid == x) || (mid+1)*(mid+1) > x {
				return mid
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// 牛顿迭代
func mySqrt(x int) int {
	y := x
	for y*y > x {
		y = (y + x/y) / 2
	}
	return y
}
