package contains_duplicate_3

func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 || k <= 0 || t < 0 {
		return false
	}

	T := int64(t)
	buckets := map[int64]int64{}
	bucketID := func(num int64) int64 {
		if num >= 0 {
			return num / (T + 1)
		}
		return (num+1)/(T+1) - 1
	}

	for i := range nums {
		num := int64(nums[i])
		if i > k {
			delete(buckets, bucketID(int64(nums[i-k-1])))
		}
		id := bucketID(num)
		if _, exists := buckets[id]; exists {
			return true
		}
		if nextBucketValue, exists := buckets[id+1]; exists {
			if nextBucketValue-num <= T {
				return true
			}
		}
		if prevBucketValue, exists := buckets[id-1]; exists {
			if num-prevBucketValue <= T {
				return true
			}
		}
		buckets[id] = num
	}
	return false
}
