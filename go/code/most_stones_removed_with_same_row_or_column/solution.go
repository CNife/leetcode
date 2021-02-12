package most_stones_removed_with_same_row_or_column

//goland:noinspection GoNilness
func RemoveStones(stones [][]int) int {
	var d disjointSet
	n := len(stones)
	result := n

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			iRoot, jRoot := d.find(i), d.find(j)
			if iRoot != jRoot && (stones[i][0] == stones[j][0] ||
				stones[i][1] == stones[j][1]) {
				d[iRoot] = jRoot
				result--
			}
		}
	}
	return n - result
}

type disjointSet []int

func (d *disjointSet) find(i int) int {
	for j := len(*d); j <= i; j++ {
		*d = append(*d, j)
	}

	if (*d)[i] != i {
		(*d)[i] = d.find((*d)[i])
	}
	return (*d)[i]
}
