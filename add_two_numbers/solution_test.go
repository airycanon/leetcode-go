package add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddTwoNumbers(t *testing.T) {
	var a *ListNode
	var b *ListNode
	var ab *ListNode

	a = NewList([]int{1, 2, 3})
	b = NewList([]int{1, 2, 3})
	ab = NewList([]int{2, 4, 6})
	assert.Equal(t, addTwoNumbers(a, b), ab)

	a = NewList([]int{2, 4, 3})
	b = NewList([]int{5, 6, 4})
	ab = NewList([]int{7, 0, 8})
	assert.Equal(t, ab, addTwoNumbers(a, b))
}

func NewList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	firstNode := &ListNode{Val: nums[0]}
	temp := firstNode
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}
	return firstNode
}
