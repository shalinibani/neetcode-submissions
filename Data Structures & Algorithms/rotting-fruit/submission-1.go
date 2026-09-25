type key struct{
	i int
	j int
}

func orangesRotting(grid [][]int) int {
    rows := len(grid)
	cols := len(grid[0])

	var q []key

	fresh := 0

	for i := 0; i<rows; i++{
		for j:=0; j<cols; j++{
			if grid[i][j] == 2{
				q = append(q, key{i,j})
			}else if grid[i][j] == 1{
				fresh++
			}
		}
	}

	time := 0

	dir := [][2]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},}

	for len(q) > 0 && fresh > 0{ 
		currSize := len(q)
		for i :=0; i<currSize; i++{
			curr := q[0]
			q = q[1:]

			for _, d := range dir{
			 ni := d[0] + curr.i
			 nj := d[1] + curr.j

			 if ni < 0 || nj < 0 || ni >=rows || nj >= cols{
				continue
			 }
			
			if grid[ni][nj] == 1{
				grid[ni][nj] = 2

				fresh--

				q = append(q, key{ni,nj})
			}
		  }
		  
		}
        time++
	}

	if fresh != 0{
		return -1
	}

	return time
}
