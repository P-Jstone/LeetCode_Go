package Search

func combine(n int, k int) [][]int {
	var res [][]int
	var track []int
	var backtrack func(startIndex, n, k int, track []int)
	backtrack = func(startIndex, n, k int, track []int) {
		if len(track) == k {
			res = append(res, append([]int(nil), track...))
		}
		if len(track)+n-startIndex+1 < k {
			return
		}
		for i := startIndex; i <= n; i++ {
			track = append(track, i)
			backtrack(i+1, n, k, track)
			track = track[:len(track)-1]
		}
	}
	backtrack(1, n, k, track)
	return res
}
