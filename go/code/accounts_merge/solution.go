package accounts_merge

import "sort"

func AccountsMerge(accounts [][]string) [][]string {
	d := newDisjointSet(len(accounts))
	emailToIndex := map[string]int{}

	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			email := account[j]
			if index, exists := emailToIndex[email]; exists {
				d.union(i, index)
			} else {
				emailToIndex[email] = i
			}
		}
	}

	indexToEmails := map[int][]string{}
	for email, index := range emailToIndex {
		index = d.find(index)
		if emails, exists := indexToEmails[index]; exists {
			indexToEmails[index] = append(emails, email)
		} else {
			indexToEmails[index] = []string{email}
		}
	}

	result := make([][]string, 0, len(indexToEmails))
	for index, emails := range indexToEmails {
		sort.Strings(emails)

		oneResult := make([]string, 0, len(emails)+1)
		oneResult = append(oneResult, accounts[index][0])
		oneResult = append(oneResult, emails...)
		result = append(result, oneResult)
	}
	return result
}

type disjointSet []int

func newDisjointSet(size int) disjointSet {
	d := make([]int, size)
	for i := range d {
		d[i] = i
	}
	return d
}

func (d disjointSet) find(i int) int {
	if d[i] != i {
		d[i] = d.find(d[i])
	}
	return d[i]
}

func (d disjointSet) union(i, j int) {
	if i != j {
		i, j = d.find(i), d.find(j)
		if i != j {
			d[i] = j
		}
	}
}
