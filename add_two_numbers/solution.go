package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	a, b, ab := l1, l2, newNode
	//进位
	carry := 0
	for nil != a || nil != b || carry > 0 {
		sum := carry
		if nil != a {
			sum += a.Val
			a = a.Next
		}
		if nil != b {
			sum += b.Val
			b = b.Next
		}

		ab.Val = sum % 10
		carry = sum / 10

		//没有进位，且两个链表都到末尾了
		if a == nil && b == nil && carry == 0 {
			return newNode
		}

		ab.Next = &ListNode{}
		ab = ab.Next
	}
	return newNode
}
