package word_break_2

func WordBreak(s string, wordList []string) []string {
	words := make(map[string]struct{})
	for _, s := range wordList {
		words[s] = struct{}{}
	}
	cache := make(map[int][]string)

	var dfs func(int) []string
	dfs = func(start int) []string {
		if cached, exists := cache[start]; exists {
			return cached
		}

		var results []string
		if start == len(s) {
			results = append(results, "")
		}
		for end := start + 1; end <= len(s); end++ {
			leftPart := s[start:end]
			if _, exists := words[leftPart]; exists {
				rightParts := dfs(end)
				for _, rightPart := range rightParts {
					oneResult := leftPart
					if len(rightPart) > 0 {
						oneResult += " " + rightPart
					}
					results = append(results, oneResult)
				}
			}
		}
		cache[start] = results
		return results
	}
	return dfs(0)
}
