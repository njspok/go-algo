package isametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion version
// complexity
// - memory with recurstion stack O(log n) for balanced tree, O(n) not balanced
// - time O(n) if tree equal structure or O(min(p,q))
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
