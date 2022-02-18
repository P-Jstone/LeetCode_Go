package GreedyAlgorithm

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	if n == 0 {
		return true
	}
	for i := 0; i < l; {
		if flowerbed[i] == 1 {
			i += 2
		} else {
			if i == l-1 {
				n--
				i++
			} else if flowerbed[i+1] == 0 {
				n--
				i += 2
			} else {
				i++
			}
		}
		if n == 0 {
			return true
		}
	}
	if n == 0 {
		return true
	} else {
		return false
	}
}
