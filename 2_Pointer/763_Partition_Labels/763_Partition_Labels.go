package GreedyAlgorithm

func partitionLabels(s string) []int {
	// 记录所有字母最后出现的位置
	var lastIndexOf [26]int
	for i, v := range s {
		lastIndexOf[v-'a'] = i
	}
	var res []int
	for l, r := 0, 0; l < len(s); l = r + 1 {
		r = lastIndexOf[s[l]-'a']
		for i := l; i < r; i++ {
			if r < lastIndexOf[s[i]-'a'] {
				r = lastIndexOf[s[i]-'a']
			}
		}
		res = append(res, r-l+1)
	}
	return res
}
