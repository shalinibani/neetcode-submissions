type MaxHeap []int

func (h MaxHeap) Len() int{
	return len(h)
}

func(h MaxHeap) Less(i,j int) bool{
	return h[i] < h[j]
}

func(h MaxHeap) Swap(i,j int){
	h[i],h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any){
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop()any{
	old := *h
	n := len(*h)
	x := old[n-1]

	*h = old[:n-1]
	return x
}

type KthLargest struct {
    k int
	h *MaxHeap
}


func Constructor(k int, nums []int) KthLargest {
	h := &MaxHeap{}
	heap.Init(h)

	for _, num := range nums{
		heap.Push(h, num)

		if h.Len() > k{
			heap.Pop(h)
		}
	}

    return KthLargest{k, h}
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.h, val)

	if this.h.Len() > this.k{
		heap.Pop(this.h)
	}

	return (*this.h)[0]
}
