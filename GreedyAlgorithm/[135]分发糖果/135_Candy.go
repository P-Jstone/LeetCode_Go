package GreedyAlgorithm

func candy(ratings []int) int {
	num := make([]int, len(ratings))
	for i, _ := range ratings {
		num[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			num[i] = num[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			if num[i-1] <= num[i] {
				num[i-1] = num[i] + 1
			}
		}
	}
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}
