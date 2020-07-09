package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := map[int]bool{
		1221:  true,
		-1221: false,
		0:     true,
		12321: true,
	}

	for key, value := range cases {
		assert.Equal(t, value, isPalindrome(key))
	}
}
