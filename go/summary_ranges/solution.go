package summary_ranges

import (
	"fmt"
	"strconv"
)

func SummaryRanges(nums []int) []string {
	n := len(nums)
	if n < 1 {
		return nil
	} else if n < 2 {
		return []string{strconv.Itoa(nums[0])}
	}

	var result []string
	start := 0
	for i := 0; i < n-1; i++ {
		if nums[i]+1 < nums[i+1] {
			if start == i {
				result = append(result, strconv.Itoa(nums[i]))
			} else {
				result = append(result, fmt.Sprintf("%v->%v", nums[start], nums[i]))
			}
			start = i + 1
		}
	}

	if start == n-1 {
		result = append(result, strconv.Itoa(nums[n-1]))
	} else {
		result = append(result, fmt.Sprintf("%v->%v", nums[start], nums[n-1]))
	}
	return result
}
