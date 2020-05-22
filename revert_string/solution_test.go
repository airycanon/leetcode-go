package revert_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RevertString(t *testing.T) {
	testCase := map[string]string{
		"hello":  "olleh",
		"Hannah": "hannaH",
	}

	for key, value := range testCase {
		result := reverseString([]byte(key))
		assert.Equal(t, value, string(result))
	}
}
