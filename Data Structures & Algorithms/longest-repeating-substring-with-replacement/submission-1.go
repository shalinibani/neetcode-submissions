func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)

	left, right := 0,0

	long := 0

	for right < len(s){
		freqMap[s[right]]++

		window := right - left + 1 

		for window - countMostFreqCharCount(freqMap) > k {
			freqMap[s[left]]--
			left++

			window = right - left +1
		}

		long = max(long, window)

		right++
	}

	return long
	
}

func max(a, b int) int{
	if a > b{
		return a
	}

	return b
}

func countMostFreqCharCount(freqMap map[byte]int) int{
	most := 0

	for _,v := range freqMap{
		if v > most{
			most = v
		}
	}

	return most
}
