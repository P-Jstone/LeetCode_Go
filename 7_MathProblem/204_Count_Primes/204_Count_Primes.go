package MP

func countPrimes(n int) int {
	// 没有比2还小的质数
	if n <= 2 {
		return 0
	}
	prime := make([]bool, n+1)
	// 除去1和数字n本身的剩余可能质数个数
	count := n - 2
	for i := 2; i <= n; i++ {
		// 如果数字 i 是质数
		if !prime[i] {
			// 遍历 i 小于 n 的倍数 j
			for j := 2 * i; j < n; j += i {
				// 如果 j 记录为质数
				if !prime[j] {
					prime[j] = true // 将 j 记录为合数
					count--
				}
			}
		}
	}
	return count
}
