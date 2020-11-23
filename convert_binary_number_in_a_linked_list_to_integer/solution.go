package convert_binary_number_in_a_linked_list_to_integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	numbers := []int{head.Val}
	next := head.Next
	for next != nil {
		numbers = append(numbers, next.Val)
		next = next.Next
	}

	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	result := 0
	for i, number := range numbers {
		if number == 1 {
			result += 1 << i
		}
	}

	return result
}
