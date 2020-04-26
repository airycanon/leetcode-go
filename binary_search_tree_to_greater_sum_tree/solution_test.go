package binary_search_tree_to_greater_sum_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BstToGst(t *testing.T) {
	input := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  7,
				Left: nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	output := &TreeNode{
		Val: 30,
		Left: &TreeNode{
			Val: 36,
			Left: &TreeNode{
				Val:   36,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  35,
				Left: nil,
				Right: &TreeNode{
					Val:   33,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 21,
			Left: &TreeNode{
				Val:   26,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  15,
				Left: nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	assert.Equal(t, bstToGst(input), output)
}
