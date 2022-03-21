package String

// 方法一中心扩展
//func countSubstrings(s string) int {
//	var res int
//	n := len(s)
//	for i := 0; i < 2*n-1; i++ {
//		l, r := i/2, i/2+i%2
//		for l >= 0 && r < n && s[l] == s[r] {
//			l--
//			r++
//			res++
//		}
//	}
//	return res
//}

// 方法二 Manacher算法
func countSubstrings(s string) int {
	n := len(s)
	t := "$#"
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}
	n = len(t)
	t += "!"

	f := make([]int, n)
	iM, rM, res := 0, 0, 0
	for i := 1; i < n; i++ {
		if i <= rM {
			f[i] = min(f[2*iM-i], rM-i+1)
		} else {
			f[i] = 1
		}

		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}

		if i+f[i]-1 > rM {
			iM = i
			rM = i + f[i] - 1
		}
		res += f[i] / 2
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
