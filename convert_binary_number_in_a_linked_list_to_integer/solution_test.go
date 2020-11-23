package convert_binary_number_in_a_linked_list_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetDecimalValue(t *testing.T) {

	cases := map[int][]int{
		5:  {1, 0, 1},
		10: {1, 0, 1, 0},
		7:  {1, 1, 1},
		31: {1, 1, 1, 1, 1},
		27: {1, 1, 0, 1, 1},
	}

	for k, v := range cases {
		assert.Equal(t, k, getDecimalValue(convert(v)))
	}
}

func convert(n []int) *ListNode {
	head := &ListNode{
		Val: n[0],
	}
	n = n[1:]

	tmp := head
	for len(n) > 0 {
		tmp.Next, n = &ListNode{
			Val: n[0],
		}, n[1:]
		tmp = tmp.Next
	}

	return head
}
