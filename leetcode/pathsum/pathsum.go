package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Complexity
// - space O(n) bad case or O(log(n)) for balanced tree
// - time O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return checkSum(root, targetSum, 0)
}

func checkSum(node *TreeNode, targetSum int, sum int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return node.Val+sum == targetSum
	}

	return checkSum(node.Left, targetSum, sum+node.Val) || checkSum(node.Right, targetSum, sum+node.Val)
}
