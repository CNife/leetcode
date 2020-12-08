package split_array_into_fibonacci_sequence

import "math"

func SplitIntoFibonacci(s string) []int {
	var result []int
	backtrack(s, &result)
	return result
}

func backtrack(s string, result *[]int) bool {
	if len(s) <= 0 && len(*result) >= 3 {
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[0] == '0' && i > 0 {
			break
		}
		num64 := string2Int64(s[:i+1])
		if num64 > math.MaxInt32 {
			break
		}
		num := int(num64)

		resLen := len(*result)
		if resLen > 1 && (*result)[resLen-1]+(*result)[resLen-2] < num {
			break
		}
		if resLen < 2 || (*result)[resLen-1]+(*result)[resLen-2] == num {
			*result = append(*result, num)
			if backtrack(s[i+1:], result) {
				return true
			}
			*result = (*result)[:len(*result)-1]
		}
	}
	return false
}

func string2Int64(s string) int64 {
	var num int64 = 0
	for i := 0; i < len(s); i++ {
		digit := s[i] - '0'
		num = num*10 + int64(digit)
	}
	return num
}
