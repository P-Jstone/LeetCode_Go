package MP

import "strconv"

func addStrings(num1 string, num2 string) string {
	c := 0
	res := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || c != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		z := x + y + c
		res = strconv.Itoa(z%10) + res
		c = z / 10
	}
	return res
}
