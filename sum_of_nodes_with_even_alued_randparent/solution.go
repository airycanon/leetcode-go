package sum_of_nodes_with_even_alued_randparent

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	var sum = 0
	var current *TreeNode
	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		left := current.Left
		right := current.Right

		if left != nil {
			queue = append(queue, left)
			if current.Val%2 == 0 {
				if left.Left != nil {
					sum += left.Left.Val
				}

				if left.Right != nil {
					sum += left.Right.Val
				}
			}
		}

		if right != nil {
			queue = append(queue, right)
			if current.Val%2 == 0 {
				if right.Left != nil {
					sum += right.Left.Val
				}

				if right.Right != nil {
					sum += right.Right.Val
				}
			}
		}
	}

	return sum
}
