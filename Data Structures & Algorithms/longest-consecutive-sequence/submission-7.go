func longestConsecutive(nums []int) int {
  numMap := make(map[int]struct{})

  for _, n := range nums{
	numMap[n] = struct{}{}
  }

  res := 0

  for i := 0; i<len(nums); i++{

	if _, ok :=numMap[nums[i]-1]; !ok{
		count := 0
		temp := nums[i]
		for {
			count++
			temp++
			if _, ok := numMap[temp]; !ok{
				break
			}
		}

		res = max(res, count) 
	}
  }

  return res
}

func max(a, b int) int{
	if a < b{
		return b
	}

	return a
}
