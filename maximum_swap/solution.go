package maximum_swap

import (
	"strconv"
)

func maximumSwap(num int) int {
	str := strconv.Itoa(num)
	if len(str) == 1 {
		return num
	}
	chars := []rune(str)

	var (
		index int
		found bool
	)

	for i := 0; i < len(chars); i++ {
		var max = chars[i]
		index = i
		for j := i + 1; j < len(chars); j++ {
			if chars[j] > chars[i] && chars[j] >= max {
				index = j
				max = chars[j]
				found = true
			}
		}

		if found {
			chars[i], chars[index] = chars[index], chars[i]
			break
		}
	}

	newNum, _ := strconv.Atoi(string(chars))

	return newNum
}
