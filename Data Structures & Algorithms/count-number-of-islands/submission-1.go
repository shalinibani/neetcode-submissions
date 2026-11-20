func numIslands(grid [][]byte) int {
    numIslands := 0

	rows := len(grid)
	cols := len(grid[0])

	var dfs func(i,j int)

	dfs = func(i,j int){
		if i < 0 || j < 0 || j >= cols || i >= rows || grid[i][j] != '1'{
			return
		}

		grid[i][j] = '#'

		dfs(i+1,j)
		dfs(i-1,j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 0; i<rows; i++{
		for j := 0; j<cols; j++{
			if grid[i][j] == '1'{
				dfs(i,j)

				numIslands++
			}
		}
	}

	return numIslands
}
