package BinarySearch

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l+((r-l)>>1)
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[l] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[l] == nums[mid] {
				l++
			}
			if nums[r] == nums[mid] {
				r--
			}
		}
	}
	return false
}
