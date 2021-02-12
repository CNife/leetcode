package how_many_numbers_are_smaller_than_the_current_number

import "sort"

func SmallerNumbersThanCurrent(nums []int) []int {
	ins := make([]idxNum, 0, len(nums))
	for idx, num := range nums {
		ins = append(ins, idxNum{idx, num})
	}
	sort.Slice(ins, func(i, j int) bool {
		return ins[i].num < ins[j].num
	})

	prevNum, prevCount := -1, -1
	for i, in := range ins {
		if in.num != prevNum {
			prevNum = in.num
			prevCount = i
		}
		ins[i] = idxNum{in.idx, prevCount}
	}
	sort.Slice(ins, func(i, j int) bool {
		return ins[i].idx < ins[j].idx
	})

	result := make([]int, len(ins))
	for i, in := range ins {
		result[i] = in.num
	}
	return result
}

type idxNum struct {
	idx, num int
}
