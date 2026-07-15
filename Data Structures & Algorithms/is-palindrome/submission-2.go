func isPalindrome(s string) bool {
 for i, j:= 0, len(s)-1; i < j; {
	if (isLetter(s[i]) || isDigit(s[i])) && (isLetter(s[j]) || isDigit(s[j])) {
		if toLower(s[i]) != toLower(s[j]){
			return false
		}

		i++
		j--
	}else{
		if !isLetter(s[i]) && !isDigit(s[i]){
			i++
		}
		if !isLetter(s[j]) && !isDigit(s[j]){
			j--
		}
	}
 }

 return true
}

func toLower(c byte) byte{
	if c >= 'A' && c<='Z'{
		return c + ('a' - 'A')
	}

	return c
}

func isLetter(c byte) bool{
	if c >= 'a' && c<='z' || (c>='A' && c<= 'Z'){
		return true
	}

	return false
}

func isDigit(c byte) bool{
	return  c >= '0' && c<='9'
}
