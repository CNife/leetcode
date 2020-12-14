package group_anagrams

func GroupAnagrams(strs []string) [][]string {
	buckets := map[int][]string{}
	for _, str := range strs {
		key := 1
		for i := 0; i < len(str); i++ {
			key *= primes[str[i]-'a']
		}
		if bucket, exists := buckets[key]; exists {
			buckets[key] = append(bucket, str)
		} else {
			buckets[key] = []string{str}
		}
	}

	result := make([][]string, 0, len(buckets))
	for _, bucket := range buckets {
		result = append(result, bucket)
	}
	return result
}

var primes = [26]int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73,
	79, 83, 89, 97, 101, 103,
}
