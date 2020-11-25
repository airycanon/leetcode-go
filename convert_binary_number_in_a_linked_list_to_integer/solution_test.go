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
	var current, head *ListNode

	for len(n) > 0 {
		node := &ListNode{Val: n[0]}
		n = n[1:]
		if current == nil {
			current = node
			head = node
		} else {
			current.Next = node
			current = node
		}
	}

	return head
}
