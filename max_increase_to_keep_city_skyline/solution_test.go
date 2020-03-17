package max_increase_to_keep_city_skyline

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxIncreaseKeepingSkyline(t *testing.T) {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}

	increment := maxIncreaseKeepingSkyline(grid)
	assert.Equal(t, increment, 35)
}
