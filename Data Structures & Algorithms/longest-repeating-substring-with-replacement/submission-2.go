func characterReplacement(s string, k int) int {
  freqMap := make(map[byte]int)

  left := 0

  long := 0
  maxFreq := 0

  for right := 0; right < len(s); right++{
	freqMap[s[right]]++

	if freqMap[s[right]] > maxFreq{
		maxFreq = freqMap[s[right]]
	}

	for (right - left + 1) - maxFreq > k{
		freqMap[s[left]]--
		left++
	}

	long = max(long, right-left+1)
  }

  return long
}
