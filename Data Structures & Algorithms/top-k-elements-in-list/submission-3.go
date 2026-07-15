type freq struct{
    num int
    freq int
}
type MaxHeap []freq

func (h MaxHeap) Len() int{
    return len(h)
}

func (h MaxHeap) Less(i,j int) bool{
 return h[i].freq>h[j].freq
}

func (h MaxHeap) Swap(i,j int){
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Pop() any{
 old := *h
 n := len(old)

 item := old[n-1]
 *h = old[:n-1]

 return item
}

func (h *MaxHeap) Push(x any){
 *h = append(*h, x.(freq))
}

func topKFrequent(nums []int, k int) []int {
 f := make(map[int]int)

 for _, n := range nums{
	f[n]++
 }

 h := &MaxHeap{}
 heap.Init(h)

	for n, fr := range f{
		heap.Push(h, freq{n,fr})
	}


 
  finalSol := make([]int, 0,k)

 for i:= 0; i<k; i++{
	finalSol = append(finalSol, heap.Pop(h).(freq).num)
 }

 return finalSol
}
