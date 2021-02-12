package jump_game_2

func Jump(nums []int) int {
	length, end, maxPos, steps := len(nums), 0, 0, 0
	for i := 0; i < length-1; i++ {
		if maxPos < i+nums[i] {
			maxPos = i + nums[i]
		}
		if i == end {
			end = maxPos
			steps++
		}
	}
	return steps
}
