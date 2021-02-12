package partition_labels

func PartitionLabels(src string) []int {
	lasts := make([]int, 26)
	for i := 0; i < len(src); i++ {
		lasts[src[i]-'a'] = i
	}

	var result []int
	start, end := 0, 0
	for i := 0; i < len(src); i++ {
		if thisEnd := lasts[src[i]-'a']; thisEnd > end {
			end = thisEnd
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
