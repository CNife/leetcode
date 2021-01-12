package sort_items_by_groups_respecting_dependencies

func SortItems(n int, m int, group []int, beforeItems [][]int) []int {
	var (
		groupItems  = make([][]int, m+n)
		groupEdges  = make([][]int, m+n)
		groupIdList = makeIndices(m + n)
		itemEdges   = make([][]int, n)
	)

	lastGroup := m
	for i, itemGroup := range group {
		if itemGroup == -1 {
			itemGroup, group[i] = lastGroup, lastGroup
			lastGroup++
		}
		groupItems[itemGroup] = append(groupItems[itemGroup], i)
	}

	var (
		itemDegree  = make([]int, n)
		groupDegree = make([]int, m+n)
	)
	for i := range beforeItems {
		curGroup := group[i]
		for _, data := range beforeItems[i] {
			dataGroup := group[data]
			if dataGroup == curGroup {
				itemDegree[i]++
				itemEdges[data] = append(itemEdges[data], i)
			} else {
				groupDegree[curGroup]++
				groupEdges[dataGroup] = append(groupEdges[dataGroup], curGroup)
			}
		}
	}

	outsideSort := toSort(groupDegree, groupEdges, groupIdList)
	if len(outsideSort) < 1 {
		return nil
	}

	result := make([]int, 0, n)
	for _, outGroup := range outsideSort {
		if len(groupItems[outGroup]) < 1 {
			continue
		}

		insideSort := toSort(itemDegree, itemEdges, groupItems[outGroup])
		if len(insideSort) < 1 {
			return nil
		}
		result = append(result, insideSort...)
	}
	return result
}

func makeIndices(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = i
	}
	return result
}

func toSort(degree []int, edges [][]int, points []int) []int {
	result := make([]int, 0, len(points))

	var queue []int
	for _, item := range points {
		if degree[item] == 0 {
			queue = append(queue, item)
		}
	}

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		list := edges[temp]
		for _, ed := range list {
			degree[ed]--
			if degree[ed] < 1 {
				queue = append(queue, ed)
			}
		}
		result = append(result, temp)
	}

	if len(result) == len(points) {
		return result
	} else {
		return nil
	}
}
