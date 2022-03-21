package String

func strStr(haystack string, needle string) int {
	m, n := len(needle), len(haystack)
	if m == 0 {
		return 0
	}
	prefix := computePrefix(needle)
	q := 0
	for i := 0; i < n; i++ {
		for q > 0 && needle[q] != haystack[i] {
			q = prefix[q-1]
		}
		if needle[q] == haystack[i] {
			q++
		}
		if q == m {
			return i - m + 1
		}
	}
	return -1
}

func computePrefix(pattern string) []int {
	P := []byte(pattern)
	m := len(pattern)
	prefix := make([]int, m)
	prefix[0] = 0
	k := 0
	for q := 1; q < m; q++ {
		for k > 0 && P[k] != P[q] {
			k = prefix[k-1]
		}
		if P[k] == P[q] {
			k++
		}
		prefix[q] = k
	}
	return prefix
}
