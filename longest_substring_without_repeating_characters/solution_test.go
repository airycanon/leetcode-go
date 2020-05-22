package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LengthOfLongestSubstring(t *testing.T) {
	case1 := "dfwefewfewf"
	assert.Equal(t, 4, lengthOfLongestSubstring(case1))

	case2 := "aaaaaa"
	assert.Equal(t, 1, lengthOfLongestSubstring(case2))

	case3 := "qwertyuiop"
	assert.Equal(t, 10, lengthOfLongestSubstring(case3))

	case4 := "qwertyuiopqwertyuiop"
	assert.Equal(t, 10, lengthOfLongestSubstring(case4))

	case6 := "pwwkew"
	assert.Equal(t, 3, lengthOfLongestSubstring(case6))
}
