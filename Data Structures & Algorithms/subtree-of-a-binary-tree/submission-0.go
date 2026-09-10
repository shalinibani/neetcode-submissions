/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil{
		return true
	}

	if (root == nil && subRoot !=nil) || (root!= nil && subRoot == nil){
		return false
	}

    visited := make(map[*TreeNode]bool)

	q := []*TreeNode{root}

	for len(q) > 0{
		n := q[0]
		q = q[1:]

		if !visited[n]{
			visited[n] = true

			if n.Val == subRoot.Val{
				balanced := isBalanced(n,subRoot)
				if balanced{
					return true
				}
			}

			if n.Left !=nil{
					q = append(q, n.Left)
				}

				if n.Right != nil{
					q = append(q, n.Right)
				}
		}
	}

	return false
}

func isBalanced(root1 *TreeNode, root2 *TreeNode)bool{
	if (root1 == nil && root2!= nil) ||(root1 != nil && root2==nil){
		return false
	}

	if root1 == nil && root2 == nil{
		return true
	}

	if root1.Val != root2.Val {
		return false
	}

	return isBalanced(root1.Left, root2.Left) && isBalanced(root1.Right, root2.Right)
}
