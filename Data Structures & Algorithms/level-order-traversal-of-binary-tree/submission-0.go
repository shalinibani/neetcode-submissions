/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type leveInfo struct{
	level int
	val int
 }



func levelOrder(root *TreeNode) [][]int {
    var res [][]int

	if root == nil{
		return res
	}

	q := []*TreeNode{root}

    for len(q) > 0{
		n := len(q)

		var temp []int
		for i:=0; i < n; i++{
			temp = append(temp, q[i].Val)

			if q[i].Left != nil{
				q = append(q, q[i].Left)
			}

			if q[i].Right != nil{
				q = append(q, q[i].Right)
			}
		}

		q = q[n:]
		res = append(res, temp)
	}

	
	return res
}
