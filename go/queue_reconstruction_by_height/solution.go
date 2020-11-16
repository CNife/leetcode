package queue_reconstruction_by_height

import "sort"

func ReconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[j][0] < people[i][0]
	})

	result := make([][]int, 0, len(people))
	for _, person := range people {
		result = insert(result, person[1], person)
	}
	return result
}

func insert(slice [][]int, index int, elem []int) [][]int {
	// 0 <= index <= len(slice)
	if index == len(slice) {
		return append(slice, elem)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = elem
	return slice
}
