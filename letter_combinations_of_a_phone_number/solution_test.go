package letter_combinations_of_a_phone_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LetterCombinations(t *testing.T) {
	results := letterCombinations("")
	assert.ElementsMatch(t, results, []string{})

	results = letterCombinations("2")
	assert.ElementsMatch(t, results, []string{"a", "b", "c"})

	results = letterCombinations("23")
	assert.ElementsMatch(t, results, []string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"})
}
