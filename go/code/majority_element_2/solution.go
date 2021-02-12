package majority_element_2

func MajorityElement(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}

	type candidate struct {
		num, count int
	}
	c1 := candidate{num: nums[0]}
	c2 := candidate{num: nums[0]}

	for _, num := range nums {
		if c1.num == num {
			c1.count++
			continue
		}
		if c2.num == num {
			c2.count++
			continue
		}
		if c1.count == 0 {
			c1.num = num
			c1.count++
			continue
		}
		if c2.count == 0 {
			c2.num = num
			c2.count++
			continue
		}
		c1.count--
		c2.count--
	}

	c1.count, c2.count = 0, 0
	for _, num := range nums {
		switch num {
		case c1.num:
			c1.count++
		case c2.num:
			c2.count++
		}
	}

	var result []int
	if c1.count > len(nums)/3 {
		result = append(result, c1.num)
	}
	if c2.count > len(nums)/3 {
		result = append(result, c2.num)
	}
	return result
}
