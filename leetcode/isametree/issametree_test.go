package isametree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name   string
		p      *TreeNode
		q      *TreeNode
		result bool
	}{
		{
			name: "nil trees",
			p:    nil, q: nil, result: true,
		},
		{
			name:   "one equal node",
			p:      &TreeNode{Val: 1},
			q:      &TreeNode{Val: 1},
			result: true,
		},
		{
			name:   "one different node",
			p:      &TreeNode{Val: 1},
			q:      &TreeNode{Val: 2},
			result: false,
		},
		{
			name: "2 level equal trees",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			result: true,
		},
		{
			name: "2 level different trees",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: nil,
			},
			q: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: &TreeNode{Val: 2},
			},
			result: false,
		},
		{
			name: "2 level different permutation trees ",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			result: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.result, isSameTree(test.p, test.q))
		})
	}
}
