package string_to_integer

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")

	if len(s) == 0 {
		return 0
	}

	var neg bool
	if s[0] == '-' || s[0] == '+' {
		neg = s[0] == '-'
		s = s[1:]
	}
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			break
		}
		n = n*10 + int(ch-'0')

		if n > math.MaxInt32 {
			if neg {
				return -math.MaxInt32 - 1
			}

			return math.MaxInt32
		}
	}

	if neg {
		n = -n
	}

	return int(n)
}
