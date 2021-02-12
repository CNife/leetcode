package regions_cut_by_slashes

func RegionsBySlashes(grid []string) int {
	n := len(grid)
	if n < 1 {
		return 0
	}

	d := newDisjointSet(n * n * 4)

	for i, line := range grid {
		for j := 0; j < len(line); j++ {

			base := 4 * (n*i + j)

			switch line[j] {
			case ' ':
				d.union(base, base+1)
				d.union(base, base+2)
				d.union(base, base+3)
			case '/':
				d.union(base, base+3)
				d.union(base+1, base+2)
			case '\\':
				d.union(base, base+1)
				d.union(base+2, base+3)
			}

			if i < n-1 {
				d.union(base+2, base+4*n)
			}
			if j < n-1 {
				d.union(base+1, base+7)
			}
		}
	}

	return d.counter
}

type disjointSet struct {
	inner   []int
	counter int
}

func newDisjointSet(size int) *disjointSet {
	inner := make([]int, size)
	for i := range inner {
		inner[i] = i
	}
	return &disjointSet{inner, size}
}

func (d *disjointSet) find(i int) int {
	parent := &d.inner[i]
	if *parent != i {
		*parent = d.find(*parent)
	}
	return *parent
}

func (d *disjointSet) union(i, j int) {
	if i != j {
		i, j = d.find(i), d.find(j)
		if i != j {
			d.inner[i] = j
			d.counter--
		}
	}
}
