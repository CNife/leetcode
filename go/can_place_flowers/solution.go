package can_place_flowers

func CanPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if n <= 0 {
		return true
	}
	if length <= 0 {
		return false
	} else if length == 1 {
		return n <= 1 && flowerbed[0] == 0
	}

	m := 0
	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		m++
	}

	for i := 1; i < length-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			m++
			if m >= n {
				return true
			}
		}
	}

	if flowerbed[length-1] == 0 && flowerbed[length-2] == 0 {
		m++
	}
	return m >= n
}
