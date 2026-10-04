func validTree(n int, edges [][]int) bool {
    graph := make([][]int, n)

	for _, e := range edges{
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	visited := make(map[int]bool)

	var dfs func(node int, parent int)bool

	dfs = func(node int, parent int) bool{
		visited[node] = true

		for _, n := range graph[node]{
			if !visited[n]{
				if !dfs(n, node){
					return false
				}
			}else if parent !=n{
				return false
			}
		}

		return true
	}

	if !dfs(0, -1){
		return false
	}

	return len(visited) == n
}
