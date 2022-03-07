package DP

func rob(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, 3)
	dp[0] = 0
	dp[1] = nums[0]
	if len(nums) < 2 {
		return dp[len(nums)]
	}
	for i := 1; i < len(nums); i++ {
		dp[2] = max(dp[1], nums[i]+dp[0])
		dp[0] = dp[1]
		dp[1] = dp[2]
	}
	return dp[2]
}
