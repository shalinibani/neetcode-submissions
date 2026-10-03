func solve(board [][]byte) {
    rows := len(board)
	cols := len(board[0])


	var dfs func (r,c int)
	dfs = func(r,c int){
		if r < 0 || c < 0 || r >=rows || c >=cols || board[r][c] != 'O'{
			return 
		}

		board[r][c] = '#'
		dfs(r+1,c)
		dfs(r-1,c)
		dfs(r,c+1)
		dfs(r,c-1)
	}

	for i := 0;i<rows;i++{
		for j := 0; j<cols ; j++{
			if board[i][j] == 'O' && (i == 0 || i == rows-1 || j == 0 || j == cols-1){
				dfs(i,j)
			}
		}
	}

	for i := 0; i<rows; i++{
		for j:= 0; j<cols; j++{
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			}else if board[i][j] == '#'{
				board[i][j] = 'O'
			}
		}
	}
}
