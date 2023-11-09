package letter_combinations_of_a_phone_number

import (
	"strconv"
	"strings"
)

func letterCombinations(digits string) []string {
	nums := strings.Split(digits, "")

	mappings := map[int64][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	results := make([]string, 0)
	for _, num := range nums {
		i, _ := strconv.ParseInt(num, 10, 64)
		if letters, exist := mappings[i]; exist {
			if len(results) == 0 {
				results = letters
				continue
			}
			newResults := make([]string, 0)
			for _, letter := range letters {
				for _, result := range results {
					newResults = append(newResults, result+letter)
				}
			}
			results = newResults
		}
	}

	return results
}
