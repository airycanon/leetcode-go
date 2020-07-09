package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	assert.Equal(t, []int{0, 1}, twoSum(nums, target))

	nums = []int{2, 11, 15, 5}
	target = 7

	assert.Equal(t, []int{0, 3}, twoSum(nums, target))
}
