package Search

func searchRange(nums []int, target int) []int {
	return []int{searchFirst(nums, target), searchLast(nums, target)}
}

func searchLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 1)
		if nums[mid] < target {
			l = mid+1
		} else if nums[mid] > target {
			r = mid-1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			l = mid+1
		}
	}
	return -1
}

func searchFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l+((r-l)>>1)
		if nums[mid] > target {
			r = mid-1
		} else if nums[mid] < target {
			l = mid+1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			r = mid-1
		}
	}
	return -1
}
