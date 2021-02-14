package couples_holding_hands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSwapsCouples(t *testing.T) {
	assert.Equal(t, 1, MinSwapsCouples([]int{0, 2, 1, 3}))
	assert.Equal(t, 0, MinSwapsCouples([]int{3, 2, 0, 1}))
}
