func numIslands(grid [][]byte) int {
    islands := 0

	rows := len(grid)
	cols := len(grid[0])

	var dfs func(r,c int)

	 dfs = func(r,c int) {
		if r >=rows || c >= cols || r < 0 || c < 0 || grid[r][c] != '1'{
			return
		}

		grid[r][c] = '0'

		dfs(r+1,c)
		dfs(r-1,c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i:= 0; i<rows;i++{
		for j:=0; j<cols; j++{
			if grid[i][j] == '1'{
				dfs(i,j)
				islands++
			}
		}
	}

	return islands


}
