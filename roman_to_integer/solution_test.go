package roman_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	cases := map[string]int{
		"III":     3,
		"IV":      4,
		"IX":      9,
		"LVIII":   58,
		"MCMXCIV": 1994,
	}

	for key, value := range cases {
		assert.Equal(t, value, romanToInt(key))
	}
}
