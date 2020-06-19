package valid_palindrome

func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for ; i < j && !isValid(s[i]); i++ {
		}
		for ; i < j && !isValid(s[j]); j-- {
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isValid(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b >= 'A' && b <= 'Z'
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
