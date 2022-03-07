package Pointer

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	n := len(numbers)
	i, j := 0, n-1
	for {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			res[0] = i + 1
			res[1] = j + 1
			return res
		}
	}
}
