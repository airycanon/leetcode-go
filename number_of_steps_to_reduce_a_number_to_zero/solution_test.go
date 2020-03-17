package number_of_steps_to_reduce_a_number_to_zero

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NumberOfSteps(t *testing.T) {
	testCase := map[int]int{
		14:  6,
		8:   4,
		123: 12,
	}

	for key, value := range testCase {
		assert.Equal(t, numberOfSteps(key), value)
	}
}
