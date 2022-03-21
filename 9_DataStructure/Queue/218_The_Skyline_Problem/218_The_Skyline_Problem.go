package DS

import (
	"container/heap"
	"sort"
)

type pair struct {
	right, height int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].height > h[j].height
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func getSkyline(buildings [][]int) [][]int {
	res := make([][]int, 0)
	n := len(buildings)
	// 将所有建筑的左右边缘坐标保存
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	// 将边缘坐标从小到大排序
	sort.Ints(boundaries)
	idx := 0
	h := hp{}
	// 从边缘数组中依次取出边缘坐标
	for _, boundary := range boundaries {
		// 将所有建筑中左边缘小于等于当前边缘坐标的建筑右边缘坐标和高度入优先队列（该队列为最大堆）
		// 此时队首元素的高度是最高的，但是右边缘坐标不一定小于当前的边缘坐标
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		// 当队列不空时，如果队首元素的右边缘小于等于当前边缘坐标，说明当前建筑不包含该坐标，其高度对此无贡献
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}
		maxn := 0
		// 队列为空说明没有建筑包含该关键点，其高度为零
		if len(h) > 0 {
			maxn = h[0].height
		}
		// 当还未选择关键点，或者当前关键点与前一个关键点的高度不一致时，直接将该关键点加入结果
		if len(res) == 0 || maxn != res[len(res)-1][1] {
			res = append(res, []int{boundary, maxn})
		}
	}
	return res
}
