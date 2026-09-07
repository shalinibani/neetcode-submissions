/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

    left := findHeight(root.Left)
	right := findHeight(root.Right)

	if abs(left-right) > 1{
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func findHeight(root *TreeNode) int{
	if root == nil{
		return 0
	}

	left := findHeight(root.Left)
	right := findHeight(root.Right)

	return 1 + max(left, right)
}

func abs(a int) int{
	if a > 0{
		return a
	}

	return -1 * a
}

func max(a, b int) int{
	if a > b{
		return a
	}

	return b
}
