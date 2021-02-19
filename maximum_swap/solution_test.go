package maximum_swap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SwapOnceToMax(t *testing.T) {
	assert.Equal(t, 1, maximumSwap(1))
	assert.Equal(t, 531, maximumSwap(135))
	assert.Equal(t, 4242, maximumSwap(2244))
	assert.Equal(t, 4422, maximumSwap(2442))
	assert.Equal(t, 4321, maximumSwap(4123))
	assert.Equal(t, 4321, maximumSwap(4312))
	assert.Equal(t, 7776, maximumSwap(7677))
	assert.Equal(t, 43321, maximumSwap(43123))
	assert.Equal(t, 4100, maximumSwap(4001))
}
