package longest_substring_without_repeating_characters

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	found := make([]rune, 0)
	chars := []rune(s)
	var (
		max   = 0
		left  = 0
		right = 0
	)

	length := len(chars)

	for left < length {
		for right < length {
			if strings.Index(string(found), string(chars[right])) > -1 {
				left++
				found = found[1:]
				break
			} else {
				found = append(found, chars[right])
				right++
			}

			if len(found) > max {
				max = len(found)
			}
		}

		if right == length {
			break
		}
	}

	return max
}
