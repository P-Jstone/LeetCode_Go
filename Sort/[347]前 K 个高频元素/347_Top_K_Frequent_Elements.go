package Sort

// 桶排序
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	maxcount := 0
	for _, v := range nums {
		count[v]++
		if count[v] > maxcount {
			maxcount = count[v]
		}
	}
	buckets := make([][]int, maxcount)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	for k, v := range count {
		buckets[v-1] = append(buckets[v-1], k)
	}
	res := make([]int, 0)
	for i := maxcount - 1; i >=0 && len(res) < k; i-- {
		if buckets[i] == nil {
			continue
		}
		for _, v := range buckets[i] {
			res = append(res, v)
		}
	}
	return res
}
