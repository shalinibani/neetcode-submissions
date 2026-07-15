func lengthOfLongestSubstring(s string) int {
 charSet := make(map[byte]bool)

 long, current := 0,0
 i,j := 0,0
 for i<len(s){
    if j<len(s) && !charSet[s[j]]{
        current++
        charSet[s[j]] = true
        j++
    }else{
        charSet = make(map[byte]bool)
        
        if current > long{
            long = current
        }

        current = 0
        i++
        j=i
    }
 }

 if current > long{
    long = current
 }

 return long
}
