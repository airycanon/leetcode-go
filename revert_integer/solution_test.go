package revert_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := map[int]int{
		123:        321,
		1534236469: 0,
		-123:       -321,
	}

	for input, excepted := range cases {
		assert.Equal(t, excepted, reverse(input))
	}
}
