func longestConsecutive(nums []int) int {
 numSet := make(map[int]struct{})

 for _, n := range nums{
	numSet[n] = struct{}{}
 }

 long := 0
 current := 0

 for _, n := range nums{
	if _, ok := numSet[n-1]; !ok{
		current = 1
		nextNum := n + 1

		for {
			if _, ok := numSet[nextNum]; ok{
				current++
				nextNum++
			}else{
				if current > long{
					long = current
				}
				break
			}
		}
	}
 }

 return long
}
