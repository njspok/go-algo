package pathsum

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {

	smallTree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	bigTree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: &TreeNode{Val: 1},
			},
		},
	}

	tests := []struct {
		name   string
		tree   *TreeNode
		sum    int
		result bool
	}{
		{name: "nil tree", tree: nil, sum: 0, result: false},

		{name: "one node", tree: &TreeNode{Val: 10}, sum: 9, result: false},
		{name: "one node", tree: &TreeNode{Val: 10}, sum: 10, result: true},

		{name: "small tree", tree: smallTree, sum: 10, result: false},
		{name: "small tree", tree: smallTree, sum: 3, result: true},
		{name: "small tree", tree: smallTree, sum: 4, result: true},

		{name: "big tree", tree: bigTree, sum: 10, result: false},
		{name: "big tree", tree: bigTree, sum: 27, result: true},
		{name: "big tree", tree: bigTree, sum: 22, result: true},
		{name: "big tree", tree: bigTree, sum: 18, result: true},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %s sum: %d", test.name, ternary(test.result, "match", "not match"), test.sum), func(t *testing.T) {
			require.Equal(t, test.result, hasPathSum(test.tree, test.sum))
		})
	}
}

func ternary[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}
