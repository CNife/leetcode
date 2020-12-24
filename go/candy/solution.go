package candy

func Candy(ratings []int) int {
	left := make([]int, len(ratings))
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := make([]int, len(ratings))
	right[len(ratings)-1] = 1
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	result := 0
	for i := range left {
		if left[i] > right[i] {
			result += left[i]
		} else {
			result += right[i]
		}
	}
	return result
}
