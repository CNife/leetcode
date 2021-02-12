package number_of_provinces

func FindCircleNum(connects [][]int) int {
	n := len(connects)
	uf, result := makeUnionFind(n), n

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if connects[i][j] == 0 {
				continue
			}

			x, y := uf.find(i), uf.find(j)
			if x != y {
				uf[x] = y
				result--
			}
		}
	}

	return result
}

type unionFind []int

func makeUnionFind(size int) unionFind {
	uf := make([]int, size)
	for i := range uf {
		uf[i] = i
	}
	return uf
}

func (uf unionFind) find(i int) int {
	if i != uf[i] {
		uf[i] = uf.find(uf[i])
	}
	return uf[i]
}
