type key struct{
	i int
	j int
}

func islandsAndTreasure(grid [][]int) {
    rows := len(grid)
	cols := len(grid[0])

	var q []key

	for i := 0; i<rows; i++{
		for j:=0; j<cols; j++{
			if grid[i][j] == 0{
				q = append(q, key{i,j})
			}	
		}
	}

	directions := [][2]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},
	}

	for len(q) > 0{
		curr := q[0]
		q = q[1:]

		for _,d := range directions{
			ni := d[0] + curr.i
			nj:= d[1] + curr.j

			if ni < 0 || nj < 0 || ni >= rows || nj >=cols{
				continue
			}

			if grid[ni][nj] != 2147483647{
				continue
			}

			grid[ni][nj] = grid[curr.i][curr.j] + 1
			q = append(q, key{ni, nj})
		}
	}
}
