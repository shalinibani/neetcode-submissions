type resIdx struct{
	i int
	j int
}

func minWindow(s string, t string) string {
	if len(s) < len(t){
		return ""
	}

    tCountMap := make(map[byte]int)
	currWindowCountMap := make(map[byte]int)

	for i:=0; i<len(t); i++{
		tCountMap[t[i]]++
	}

	have, need := 0, len(tCountMap)

	res, resLen := resIdx{-1,-1}, math.MaxInt

	l,r := 0, 0

	for r<len(s){
		currWindowCountMap[s[r]]++

		if _, ok :=tCountMap[s[r]]; ok && currWindowCountMap[s[r]] == tCountMap[s[r]]{
			have++
		}

		for have == need{
			if(r-l+1) < resLen{
				res, resLen = resIdx{l,r}, (r-l+1)
			}

			currWindowCountMap[s[l]]--
			
			if _, ok := tCountMap[s[l]]; ok && currWindowCountMap[s[l]] < tCountMap[s[l]]{
				have--
			}

			l++
		}

		r++
	}

	if resLen != math.MaxInt{
		return s[res.i:res.j+1]
	}

	return ""
}
