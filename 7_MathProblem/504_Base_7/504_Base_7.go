package MP

func convertToBase7(num int) string {
	var convert func(num int) string
	convert = func(num int) string {
		res := make([]byte, 0)
		for num != 0 {
			n := num % 7
			res = append(res, '0'+byte(n))
			num /= 7
		}
		l, r := 0, len(res)-1
		for l < r {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		}
		return string(res)
	}
	var res string
	if num == 0 {
		res = "0"
	} else if num < 0 {
		res = "-" + convert(-num)
	} else {
		res = convert(num)
	}
	return res
}

func main() {
	res := convertToBase7(100)
	println(res)
}
