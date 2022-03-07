package Pointer

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	sFreq, tFreq := make(map[uint8]int), make(map[uint8]int)
	var res string
	// 最后结果的窗口指针一开始应该不在数组中，故而设置为-1
	// 此处窗口右边界设为 -1 实际对应的元素为 r+1 这是因为有时增加了窗口边界却没有提供记录相关字符的频数
	l, r, lFin, rFin, minW, cnt := 0, -1, -1, -1, len(s) + 1, 0
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}
	for l < len(s) - len(t) + 1 {
		// 当前窗口没有到达字符串 s 的末尾并且者未包含字符串 t 中的全部字符则将滑动窗口向右移动
		if r+1 < len(s) && cnt < len(t) {
			sFreq[s[r+1]]++
			if sFreq[s[r+1]] <= tFreq[s[r+1]] {
				cnt++
			}
			r++
		} else {
			// 当前窗口中包含字符串 t 中的所有字符，并且窗口大小也是目前最小的
			// 记录当前窗口位置，更改当前最小窗口记录
			if r-l+1 < minW && cnt == len(t) {
				minW = r-l+1
				lFin = l
				rFin = r
			}
			if sFreq[s[l]] == tFreq[s[l]] {
				cnt--
			}
			sFreq[s[l]]--
			l++
		}
	}
	if lFin != -1 {
		res = s[lFin:rFin+1]
	}
	return res
}
