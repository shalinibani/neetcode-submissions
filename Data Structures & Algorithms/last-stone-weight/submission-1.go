type MaxHeap []int

func (h *MaxHeap) Push(x any){
   *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any{
    n := len(*h)
    old := *h

    x := old[n-1]
    *h = old[:n-1]

    return x
}

func (h MaxHeap) Swap(i,j int) {
    h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i,j int) bool{
    return h[i] > h[j]
}

func (h MaxHeap) Len() int{
    return len(h)
}

func lastStoneWeight(stones []int) int {
    h := MaxHeap(stones)

    heap.Init(&h)

    for h.Len() > 1{
        s1 := heap.Pop(&h).(int)
        s2 := heap.Pop(&h).(int)

        if s1!=s2{
            heap.Push(&h, s1-s2)
        }
    }

    if h.Len() > 0{
        return heap.Pop(&h).(int)
    }

    return 0
}
