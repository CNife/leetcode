package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, num := range nums {
		if _, exists := set[num]; exists {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
