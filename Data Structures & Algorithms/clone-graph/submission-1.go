/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil{
		return nil
	}

	graphMap := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node{
		if n, ok := graphMap[node]; ok{
			return n
		}

		cpNode := &Node{Val: node.Val}
		graphMap[node] = cpNode

		for i:=0; i<len(node.Neighbors); i++{
			cpNode.Neighbors = append(cpNode.Neighbors, dfs(node.Neighbors[i]))
		}

		return cpNode
	}

	return dfs(node)
}
