package sum_of_nodes_with_even_valued_grandparent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//[6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]

func Test_SumEvenGrandparent(t *testing.T) {
	root := TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}

	root.Left = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	root.Right = &TreeNode{
		Val: 8,
	}

	root.Left.Left = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 9,
		},
	}

	root.Left.Right = &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	root.Right.Left = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	root.Right.Right = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
		},
	}

	assert.Equal(t, sumEvenGrandparent(&root), 18)
}
