package DP

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, 2)
	dp[0] = 1
	dp[1] = 2
	for i := 3; i <= n; i++ {
		sum := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}
