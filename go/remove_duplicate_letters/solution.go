package remove_duplicate_letters

func RemoveDuplicateLetters(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if find(stack, s[i]) {
			continue
		}
		for len(stack) > 0 {
			stackTop := stack[len(stack)-1]
			if stackTop > s[i] && stringContainsByte(s, stackTop, i) {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}

func find(slice []byte, x byte) bool {
	for _, b := range slice {
		if b == x {
			return true
		}
	}
	return false
}

func stringContainsByte(s string, b byte, start int) bool {
	for i := start; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}
