func topKFrequent(nums []int, k int) []int {
 freq := make(map[int]int)

 for _, n := range nums{
	freq[n]++
 }

 type f struct{
	num int
	freq int
 }

 sortedFreq := make([]f,0, len(freq))

	for n, fr := range freq{
		sortedFreq = append(sortedFreq, f{n,fr})
	}

 sort.Slice(sortedFreq, func(i,j int)bool{
	if (sortedFreq[i].freq>sortedFreq[j].freq){
		return true
	}

	return false
 })



 finalSol := make([]int, 0,k)

 for i:= 0; i<k; i++{
	finalSol = append(finalSol, sortedFreq[i].num)
 }

 return finalSol
}
