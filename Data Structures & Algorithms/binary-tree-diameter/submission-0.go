/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDia := 0
	findHeight(root, &maxDia)

	return maxDia
}

func findHeight(root *TreeNode, maxDia *int) int{
	if root == nil{
		return 0
	}

	left := findHeight(root.Left, maxDia)
	right := findHeight(root.Right, maxDia)

	*maxDia = max(*maxDia, left + right)

	return 1 + max(left,right)
}

func max(a, b int)int{
	if a>b{
		return a
	}

	return b
}
