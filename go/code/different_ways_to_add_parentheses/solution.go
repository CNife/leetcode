package different_ways_to_add_parentheses

import "strconv"

func DiffWaysToCompute(input string) []int {
	if n, err := strconv.Atoi(input); err == nil {
		return []int{n}
	}

	var result []int
	for i := 0; i < len(input); i++ {
		ch := input[i]
		if isOperator(ch) {
			left := DiffWaysToCompute(input[:i])
			right := DiffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch ch {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}
	return result
}

func isOperator(ch byte) bool {
	switch ch {
	case '+', '-', '*':
		return true
	default:
		return false
	}
}
