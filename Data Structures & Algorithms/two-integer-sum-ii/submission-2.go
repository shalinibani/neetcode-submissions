func twoSum(numbers []int, target int) []int {
 numMap := make(map[int][]int)

 for i,n := range numbers{
	numMap[n] = append(numMap[n], i)
 }

 for idx1, n := range numbers{
	rem := target - n

	if idx2, ok := numMap[rem]; ok {
		if idx2[0]!=idx1{
			return []int{idx1+1, idx2[0]+1}
		}
		return []int{idx1+1, idx2[1]+1}
	}
 }

 return []int{}
}
