package Search

func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	used := make([]bool, len(nums))
	var backtrack func(nums, track []int, used []bool)
	backtrack = func(nums, track []int, used []bool) {
		if len(track) == len(nums) {
			res = append(res, append([]int(nil), track...))
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			track = append(track, nums[i])
			backtrack(nums, track, used)
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrack(nums, track, used)
	return res
}
