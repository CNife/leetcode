package pascals_triangle_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var pascalsTriangle = [][]int{
	{1},
	{1, 1},
	{1, 2, 1},
	{1, 3, 3, 1},
	{1, 4, 6, 4, 1},
	{1, 5, 10, 10, 5, 1},
	{1, 6, 15, 20, 15, 6, 1},
	{1, 7, 21, 35, 35, 21, 7, 1},
	{1, 8, 28, 56, 70, 56, 28, 8, 1},
}

func TestGetRow(t *testing.T) {
	for i, row := range pascalsTriangle {
		assert.Equal(t, row, GetRow(i))
	}
}
