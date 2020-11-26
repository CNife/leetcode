package maximum_gap

import (
	"math/bits"
)

func MaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	radixSort(nums)
	maxGap := 0
	for i := 1; i < len(nums); i++ {
		if gap := nums[i] - nums[i-1]; maxGap < gap {
			maxGap = gap
		}
	}
	return maxGap
}

func radixSort(nums []int) {
	for i := 0; i < bits.UintSize/4; i++ {
		var buckets [16][]int
		finished := true
		for _, num := range nums {
			shiftNum := num >> (i * 4)
			if shiftNum != 0 {
				finished = false
			}
			index := (num >> (i * 4)) & 0xF
			buckets[index] = append(buckets[index], num)
		}
		if finished {
			break
		}

		k := 0
		for j := 0; j < 16; j++ {
			for _, num := range buckets[j] {
				nums[k] = num
				k++
			}
		}
	}
}
