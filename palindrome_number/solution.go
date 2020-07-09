package palindrome_number

import "strconv"

func isPalindrome(x int) bool {
	chars := []rune(strconv.Itoa(x))

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}

	return true
}
