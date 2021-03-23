package three_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, 0, 1}, {-1, -1, 2}}

	assert.Equal(t, expected, threeSum(nums))
}
