package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
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

func (node *ListNode) Equals(m *ListNode) bool {
	n := node
	for n != nil || m != nil {
		if (n == nil && m != nil) || (n != nil && m == nil) {
			return false
		}

		if n.Val != m.Val {
			return false
		}

		n = n.Next
		m = m.Next
	}

	return true
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
