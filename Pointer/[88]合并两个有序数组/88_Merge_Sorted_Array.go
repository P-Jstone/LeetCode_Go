package Pointer

func merge(nums1 []int, m int, nums2 []int, n int) {
	for pos := m + n; m > 0 && n > 0; pos-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[pos-1] = nums1[m-1]
			m--
		} else {
			nums1[pos-1] = nums2[n-1]
			n--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
