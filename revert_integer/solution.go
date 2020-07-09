package revert_integer

import (
	"math"
)

func reverse(x int) int {
	isPositive := x > 0
	if !isPositive {
		x = int(math.Abs(float64(x)))
	}

	result := 0
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}

	if !isPositive {
		result = -result
	}

	if result < -math.MaxInt32 || result > math.MaxInt32 {
		return 0
	}

	return result
}
