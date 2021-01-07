package evaluate_division

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := graph{}
	for i := range equations {
		g.insert(equations[i][0], equations[i][1], values[i])
	}

	result := make([]float64, len(queries))
	for i, query := range queries {
		result[i] = g.calc(query[0], query[1])
	}
	return result
}

type graph map[string][]graphItem

type graphItem struct {
	divisor string
	value   float64
}

func (g graph) insert(dividend, divisor string, value float64) {
	g.insertImpl(dividend, divisor, value)
	g.insertImpl(divisor, dividend, 1/value)
}

func (g graph) insertImpl(dividend, divisor string, value float64) {
	if s, exists := g[dividend]; exists {
		g[dividend] = append(s, graphItem{divisor, value})
	} else {
		g[dividend] = []graphItem{{divisor, value}}
	}
}

func (g graph) calc(dividend, divisor string) float64 {
	if !g.containsDividend(dividend) || !g.containsDividend(divisor) {
		return -1.0
	}

	visited := map[string]struct{}{}
	queue := []graphItem{{dividend, 1.0}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		for _, nextItem := range g[item.divisor] {
			if nextItem.divisor == divisor {
				return nextItem.value * item.value
			}
			if _, exists := visited[nextItem.divisor]; !exists {
				visited[nextItem.divisor] = struct{}{}
				queue = append(queue, graphItem{nextItem.divisor, nextItem.value * item.value})
			}
		}
	}
	return -1.0
}

func (g graph) containsDividend(dividend string) bool {
	_, exists := g[dividend]
	return exists
}
