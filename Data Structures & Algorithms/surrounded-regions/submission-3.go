func solve(board [][]byte) {
    rows := len(board)
	cols := len(board[0])

    var dfs func(i,j int)
	dfs = func(i, j int){
		if i < 0 || j < 0 || i >= rows || j >= cols || board[i][j]!= 'O'{
			return
		}

		board[i][j] = '#'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 0; i<rows; i++{
		for j := 0 ; j<cols; j++{
			if board[i][j] == 'O' && ((i == 0 || i == rows -1) ||( j == 0 || j == cols -1)){
				dfs(i,j)
			}
		}
	}

	for i := 0; i<rows; i++{
		for j := 0; j<cols; j++{
			if board[i][j] == '#'{
				board[i][j] = 'O'
			}else if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
		}
	}
}
