type MaxHeap []int

func (h MaxHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Len() int{
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool{
	return h[i] > h[j]
}

func (h *MaxHeap) Push(x any){
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any{
	old := *h
	n := len(*h)

	x := old[n-1]
	*h = old[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
   freqMap := make(map[byte]int)

   for _, t := range tasks{
	freqMap[t]++
   }

   h := &MaxHeap{}
   heap.Init(h)

   for _,v := range freqMap{
	heap.Push(h, v)
   }

   time := 0

	cooldownQ := [][2]int{}
   for h.Len() > 0 || len(cooldownQ) != 0{
	 time++

	 if len(cooldownQ) > 0 && cooldownQ[0][1] == time{
		n := cooldownQ[0]
		cooldownQ = cooldownQ[1:]
		heap.Push(h, n[0])
	 }

	 if h.Len() > 0{
		count := heap.Pop(h).(int) - 1

		if count > 0{
			cooldownQ = append(cooldownQ, [2]int{count, time+n+1})
		}
	 }
   }

   return time
}
