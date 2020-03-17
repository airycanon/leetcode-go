package revert_string

func reverseString(s []byte) []byte {
	var i, j = 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}
