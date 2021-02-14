package kth_largest_element_in_a_stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthLargest(t *testing.T) {
	l := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, l.Add(3))
	assert.Equal(t, 5, l.Add(5))
	assert.Equal(t, 5, l.Add(10))
	assert.Equal(t, 8, l.Add(9))
	assert.Equal(t, 8, l.Add(4))
}
