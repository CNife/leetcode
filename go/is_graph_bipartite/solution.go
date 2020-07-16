package is_graph_bipartite

func IsBipartite(graph [][]int) bool {
	colors := make([]Color, len(graph))
	for start := 0; start >= 0; start = nextNoneColor(colors) {
		queue := []int{start}
		color := RED
		for qLen := 1; qLen > 0; qLen = len(queue) {
			for i := 0; i < qLen; i++ {
				node := queue[0]
				queue = queue[1:]

				colors[node] = color
				for _, neighbor := range graph[node] {
					neighborColor := colors[neighbor]
					if neighborColor == NONE {
						queue = append(queue, neighbor)
					} else if neighborColor == color {
						return false
					}
				}
			}
			color = ^color
		}
	}
	return true
}

func nextNoneColor(colors []Color) int {
	for i, color := range colors {
		if color == NONE {
			return i
		}
	}
	return -1
}

type Color uint8

//noinspection GoUnusedConst
const (
	NONE  Color = 0x00
	RED   Color = 0x0F
	BLACK Color = 0xF0
)
