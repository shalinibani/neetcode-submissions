func checkInclusion(s1 string, s2 string) bool {
  freqMapS1 := make(map[byte]int)
  freqMapS2 := make(map[byte]int)

  for i:=0; i<len(s1); i++{
	freqMapS1[s1[i]]++
  }

  left, right := 0,0

  for right < len(s2){
    if _, ok := freqMapS1[s2[right]];ok{
		j := 0
		for j < len(s1) && right < len(s2){
			freqMapS2[s2[right]]++
			j++
			right++
		}
	}
		left++
		right = left

		isPerm := true
		for num, freq := range freqMapS1{
			if freqMapS2[num] != freq{
				isPerm = false
				break
			}
		}

		if isPerm{
			return true
		}

	freqMapS2 = make(map[byte]int)
  }

  return false
}
