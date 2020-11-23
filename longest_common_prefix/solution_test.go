package longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LongestCommonPrefix(t *testing.T) {

	type TestCase struct {
		Input  []string
		Output string
	}

	cases := []TestCase{
		{
			Input:  []string{"flower", "flow", "flight"},
			Output: "fl",
		},
		{
			Input:  []string{"flower", "flow", "flowgwerew"},
			Output: "flow",
		},
		{
			Input:  []string{"qfe", "frgerlow", "ergqwe"},
			Output: "",
		},
		{
			Input:  []string{"flower", "flower", "flower"},
			Output: "flower",
		},
		{
			Input:  []string{"", "agwefrew343434", "agwefrewhretre"},
			Output: "",
		},
	}

	for _, item := range cases {
		assert.Equal(t, item.Output, longestCommonPrefix(item.Input))
	}
}
