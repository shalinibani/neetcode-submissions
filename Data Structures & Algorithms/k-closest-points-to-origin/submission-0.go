type Point struct{
	x,y int
	dist int
}

type MinHeap []Point

func (h *MinHeap) Push(x any){
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop()any{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func (h MinHeap) Len() int{
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool{
	return h[i].dist < h[j].dist
}

func kClosest(points [][]int, k int) [][]int {
	h := &MinHeap{}
	heap.Init(h)
	for _, point := range points{
		x, y := point[0], point[1]
		
		dist := x*x + y*y
		heap.Push(h, Point{x,y,dist})
	}

	finalRes := make([][]int, 0, k)

	for i:= 0;i < k; i++{
		p := heap.Pop(h).(Point)
		finalRes = append(finalRes, []int{p.x, p.y})
	}

	return finalRes
}
