func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0

	rows := len(grid)
	cols := len(grid[0])

	var dfs func (r, c int, area *int)
	dfs = func (r, c int, area *int){
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0{
			return
		}

		grid[r][c] = 0
		*area++

		dfs(r,c+1,area)
		dfs(r,c-1,area)
		dfs(r+1, c,area)
		dfs(r-1,c,area)
	}

	for i := 0 ; i<rows; i++{
		for j := 0 ; j<cols ; j++{
			var area int

			if grid[i][j] == 1{
				dfs(i,j,&area)

				maxArea = max(area, maxArea) 
			}
		}
	}

	return maxArea
}

func max(a,b int) int{
	if a > b{
		return a
	}

	return b
}
