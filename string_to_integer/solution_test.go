package string_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	assert.Equal(t, -42, myAtoi("    -42"))
	assert.Equal(t, -2147483648, myAtoi("    -91283472332"))
	assert.Equal(t, 0, myAtoi("+-12"))
	assert.Equal(t, 2147483647, myAtoi("9223372036854775808"))
	assert.Equal(t, 12345678, myAtoi("   0000000000012345678"))
	assert.Equal(t, 0, myAtoi("00000-42a1234"))
}
