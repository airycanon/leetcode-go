package binary_search_tree_to_greater_sum_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	nodes := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)
	var current = root

	//先按照中序遍历二叉树
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		if len(stack) > 0 {
			last := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			nodes = append(nodes, last)
			current = last.Right
		}
	}

	//逆向输出并累加
	for i := len(nodes) - 1; i >= 0; i-- {
		if i < len(nodes)-1 {
			nodes[i].Val += nodes[i+1].Val
		}
	}

	return root
}
