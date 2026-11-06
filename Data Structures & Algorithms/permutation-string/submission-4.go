func checkInclusion(s1 string, s2 string) bool {
  if len(s1) > len(s2){
	return false
  }

  freqMapS1 := make(map[byte]int)
  freqMapS2 := make(map[byte]int)

  left,right := 0,0

  for i := 0; i<len(s1) ; i++{
	freqMapS1[s1[i]]++
	freqMapS2[s2[i]]++
  }

  if mapsEqual(freqMapS1, freqMapS2){
	return true
  }

  for right = len(s1); right < len(s2) ; right++{
	freqMapS2[s2[right]]++
	freqMapS2[s2[left]]--

	if freqMapS2[s2[left]] == 0{
		delete(freqMapS2, s2[left])
	}

	left++

	if mapsEqual(freqMapS1, freqMapS2){
		return true
	}

  }

  return false
}

func mapsEqual(a,b map[byte]int) bool{
	if len(a) != len(b){
		return false
	}

	for k,v := range a{
		if b[k] != v{
			return false
		}
	}

	return true
}
