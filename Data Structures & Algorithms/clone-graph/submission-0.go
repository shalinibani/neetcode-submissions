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
    
	graph2 := make(map[*Node]*Node)

	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node{
		if n, ok := graph2[node]; ok{
			return n
		}
		
		copyNode := &Node{Val: node.Val}
		graph2[node] = copyNode

		for i := 0 ; i<len(node.Neighbors); i++{
			graph2[node].Neighbors = append(graph2[node].Neighbors, dfs(node.Neighbors[i]))
		}

		return copyNode
	}

	return dfs(node)
}
