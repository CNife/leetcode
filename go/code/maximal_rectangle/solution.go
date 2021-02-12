package maximal_rectangle

func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return 0
	}

	maxArea := 0
	dp := make([]int, len(matrix[0]))

	for _, row := range matrix {
		for i, ch := range row {
			if ch == '1' {
				dp[i]++
			} else {
				dp[i] = 0
			}
		}
		maxArea = max(maxArea, maxRectangleInHistogram(dp))
	}

	return maxArea
}

func maxRectangleInHistogram(heights []int) int {
	var s stack = []int{-1}
	maxArea := 0

	for i, height := range heights {
		for {
			top := s.top()
			if top != -1 && heights[top] >= height {
				s.pop()
				newTop := s.top()
				area := heights[top] * (i - newTop - 1)
				maxArea = max(maxArea, area)
			} else {
				break
			}
		}
		s.push(i)
	}

	for {
		top := s.top()
		s.pop()
		if top != -1 {
			newTop := s.top()
			area := heights[top] * (len(heights) - newTop - 1)
			maxArea = max(maxArea, area)
		} else {
			break
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s stack) top() int {
	return s[len(s)-1]
}
