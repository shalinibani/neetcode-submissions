func isPalindrome(s string) bool {
 processed := make([]rune,0,len(s))

 for _, r := range s{
	if unicode.IsLetter(r) || unicode.IsDigit(r){
		processed = append(processed, unicode.ToLower(r))
	}
 }
 for i, j:= 0, len(processed)-1; i<len(processed)/2; i,j = i+1, j-1{
	if processed[i] != processed[j]{
		return false
	}
 }

 return true
}
