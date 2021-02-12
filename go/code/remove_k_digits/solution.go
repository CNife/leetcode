package remove_k_digits

import "strings"

func RemoveKDigits(num string, k int) string {
	var stack []byte
	remain := len(num) - k
	for i := 0; i < len(num); i++ {
		digit := num[i]
		for k != 0 && len(stack) > 0 && stack[len(stack)-1] > digit {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}

	result := strings.TrimLeft(string(stack[:remain]), "0")
	if len(result) > 0 {
		return result
	} else {
		return "0"
	}
}
