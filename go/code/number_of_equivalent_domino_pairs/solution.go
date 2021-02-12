package number_of_equivalent_domino_pairs

import "sort"

func NumEquivDominoPairs(dominoes [][]int) int {
	n := len(dominoes)
	if n < 2 {
		return 0
	}

	for _, domino := range dominoes {
		if domino[0] > domino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}
	}
	sort.Slice(dominoes, func(i, j int) bool {
		if dominoes[i][0] == dominoes[j][0] {
			return dominoes[i][1] < dominoes[j][1]
		}
		return dominoes[i][0] < dominoes[j][0]
	})

	var result, prevCount int
	var prev []int
	for _, domino := range dominoes {
		if prev != nil {
			if domino[0] == prev[0] && domino[1] == prev[1] {
				prevCount++
				result += prevCount
				continue
			} else {
				prevCount = 0
			}
		}
		prev = domino
	}

	return result
}
