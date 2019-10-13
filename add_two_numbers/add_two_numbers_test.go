package add_two_numbers

import "testing"

func TestListNode_Equals(t *testing.T) {
	var a *ListNode
	var b *ListNode

	a = NewList([]int{1, 2, 3})
	b = NewList([]int{1, 2, 3})
	if a.Equals(b) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	a = NewList([]int{1, 2, 3})
	b = NewList([]int{1, 2, 4})
	if !a.Equals(b) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	a = NewList([]int{1, 2, 3})
	b = NewList([]int{1, 2, 3, 4})
	if !a.Equals(b) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}

func Test_AddTwoNumbers(t *testing.T) {
	var a *ListNode
	var b *ListNode
	var ab *ListNode

	a = NewList([]int{1, 2, 3})
	b = NewList([]int{1, 2, 3})
	ab = NewList([]int{2, 4, 6})
	if addTwoNumbers(a, b).Equals(ab) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}

	a = NewList([]int{2, 4, 3})
	b = NewList([]int{5, 6, 4})
	ab = NewList([]int{7, 0, 8})

	if addTwoNumbers(a, b).Equals(ab) {
		t.Log("Pass")
	} else {
		t.Error("Fail")
	}
}
