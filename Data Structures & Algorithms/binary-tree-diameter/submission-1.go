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

	var findHeight func (root *TreeNode)int

	findHeight = func(root *TreeNode) int{
		if root == nil{
			return 0
		}

		left := findHeight(root.Left)
		right := findHeight(root.Right)

		maxDia = max(left+right, maxDia)

		return 1+max(left,right)
	}

	findHeight(root)

	return maxDia 
}
