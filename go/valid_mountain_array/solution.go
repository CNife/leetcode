package valid_mountain_array

const (
	START = iota
	UP
	DOWN
)

func ValidMountainArray(array []int) bool {
	n := len(array)
	if n < 3 {
		return false
	}

	state := START
	for i := 1; i < n; i++ {
		left, right := array[i-1], array[i]
		if left == right || state == START && left > right || state == DOWN && left < right {
			return false
		}
		if left > right {
			state = DOWN
		} else {
			state = UP
		}
	}
	return state == DOWN
}
