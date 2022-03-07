package Sort

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	// 必须添加该rand函数，否则无法使用randomPartition
	rand.Seed(time.Now().UnixNano())
	pos := findKthLargestPosition(nums, 0, len(nums)-1, len(nums)-k)
	return nums[pos]
}

func findKthLargestPosition(nums []int, l, r, k int) int {
	q := randomPartition(nums, l, r)
	if q == k {
		return q
	} else if q < k {
		return findKthLargestPosition(nums, q+1, r, k)
	} else {
		return findKthLargestPosition(nums, l, q-1, k)
	}
}

func randomPartition(nums []int, l, r int) int {
	i := rand.Int() % (r - l + 1) + l
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	i := l-1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i+1
}
