package main

func isIsomorphic(s string, t string) bool {
	ms, mt := make(map[byte]int, 0), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if ms[s[i]] != mt[t[i]] {
			return false
		}
		// 由于初始化为零，所以如果当前两个字符一个第一次出现，一个上一次出现的位置为0，则会判断两者符合条件，故而需要在当前的位置加一
		ms[s[i]] = i + 1
		mt[t[i]] = i + 1
	}
	return true
}
