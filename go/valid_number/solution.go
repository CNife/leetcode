package valid_number

func IsNumber(s string) bool {
	state := 0
	for i := 0; i < len(s); i++ {
		state = table[state][charState(s[i])]
		if state < 0 {
			return false
		}
	}
	return result[state]
}

var table = [9][6]int{
	{0, 1, 6, 2, -1, -1},
	{-1, -1, 6, 2, -1, -1},
	{-1, -1, 3, -1, -1, -1},
	{8, -1, 3, -1, 4, -1},
	{-1, 7, 5, -1, -1, -1},
	{8, -1, 5, -1, -1, -1},
	{8, -1, 6, 3, 4, -1},
	{-1, -1, 5, -1, -1, -1},
	{8, -1, -1, -1, -1, -1},
}

var result = [9]bool{
	false, false, false, true, false, true, true, false, true,
}

func charState(c byte) int {
	switch c {
	case ' ':
		return 0
	case '+', '-':
		return 1
	case '.':
		return 3
	case 'e':
		return 4
	default:
		if c >= '0' && c <= '9' {
			return 2
		} else {
			return 5
		}
	}
}
