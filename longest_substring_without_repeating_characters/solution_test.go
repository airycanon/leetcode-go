package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LengthOfLongestSubstring(t *testing.T) {
	case1 := "dfwefewfewf"
	assert.Equal(t, lengthOfLongestSubstring(case1), 4)

	case2 := "aaaaaa"
	assert.Equal(t, lengthOfLongestSubstring(case2), 1)

	case3 := "qwertyuiop"
	assert.Equal(t, lengthOfLongestSubstring(case3), 10)

	case4 := "qwertyuiopqwertyuiop"
	assert.Equal(t, lengthOfLongestSubstring(case4), 10)

	case5 := "qwertyuiopnqwertyuiop"
	assert.Equal(t, lengthOfLongestSubstring(case5), 11)
}
