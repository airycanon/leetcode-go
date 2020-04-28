package longest_substring_without_repeating_characters

import "strings"

func lengthOfLongestSubstring(s string) int {
	found := make([]rune, 0)
	chars := []rune(s)
	var (
		first = 0
		last  = 0
	)
	for first < len(chars) {
		if strings.Contains(string(found), string(chars[first])) {
			last++
		} else {
			found = append(found, chars[first])
		}
		first++
	}

	return len(found)
}
