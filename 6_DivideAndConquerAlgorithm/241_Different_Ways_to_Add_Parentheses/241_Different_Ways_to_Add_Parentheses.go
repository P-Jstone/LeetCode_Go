package DCA

import (
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	res := make([]int, 0)
	if number, err := strconv.Atoi(expression); err == nil {
		res = append(res, number)
		return res
	}
	for i, c := range expression {
		ch := string(c)
		if ch == "+" || ch == "-" || ch == "*" {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, leftNum := range left {
				for _, rightNum := range right {
					var num int
					if ch == "+" {
						num = leftNum + rightNum
					} else if ch == "-" {
						num = leftNum - rightNum
					} else {
						num = leftNum * rightNum
					}
					res = append(res, num)
				}
			}
		}
	}
	return res
}
