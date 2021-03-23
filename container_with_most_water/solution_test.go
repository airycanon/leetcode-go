package container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 16, maxArea([]int{4, 3, 2, 1, 4}))
}
