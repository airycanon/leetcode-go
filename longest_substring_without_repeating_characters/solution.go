package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	found := make(map[rune]int)
	chars := []rune(s)
	var (
		max   = 0
		left  = 0
		right = 0
	)

	length := len(chars)

	for left < length {
		for right < length {
			if _, exist := found[chars[right]]; exist {
				delete(found, chars[left])
				left++
				break
			} else {
				found[chars[right]] = 1
				right++
			}

			if right-left > max {
				max = right - left
			}
		}

		if right == length {
			break
		}
	}

	return max
}
