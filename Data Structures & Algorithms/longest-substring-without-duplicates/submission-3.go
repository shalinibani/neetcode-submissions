func lengthOfLongestSubstring(s string) int {
 charSet := make(map[byte]bool)

 left := 0
 long := 0

 for right := 0; right < len(s) ; right++{
    for charSet[s[right]]{
        charSet[s[left]] = false
        left++
    }

    charSet[s[right]] = true

    if right - left + 1 > long{
        long = right - left + 1
    }
 }

 return long
}
