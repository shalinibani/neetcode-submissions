/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil{
		return true
	}

	var findHeight func (root *TreeNode)(int, bool)

	findHeight = func(root *TreeNode) (int, bool){
		if root == nil{
			return 0,true
		}

		left, balancedLeft := findHeight(root.Left)
		right, balancedRight := findHeight(root.Right)

		if !balancedLeft || !balancedRight || abs(left , right) > 1{
			return 0, false
		}

		return 1+max(left,right), true
	}

	_, balanced := findHeight(root)

	return balanced

}

func abs(a,b int)int{
	if a > b{
		return a-b
	}

	return b-a
}
