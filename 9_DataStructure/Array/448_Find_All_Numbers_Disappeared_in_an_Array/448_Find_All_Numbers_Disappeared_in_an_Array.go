package DS

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for k, v := range nums {
		if v <= n {
			res = append(res, k+1)
		}
	}
	return res
}
