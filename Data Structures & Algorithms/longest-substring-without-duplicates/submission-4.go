func lengthOfLongestSubstring(s string) int {
	freq := make(map[byte]bool)
	maxLen := 0

	left := 0

	for right := 0; right < len(s); right++{
		for freq[s[right]]{
			freq[s[left]] = false
			left++
		}
		freq[s[right]] = true
		

		maxLen = max(maxLen, right - left + 1)
	}

	return maxLen
}
