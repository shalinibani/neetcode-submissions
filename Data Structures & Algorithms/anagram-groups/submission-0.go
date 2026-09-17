type key [26]int
func groupAnagrams(strs []string) [][]string {
 mapAna := make(map[key][]string)

 for i := 0; i <len(strs); i++{
	var alph [26]int 

	for j := 0; j< len(strs[i]);j++{
		alph[strs[i][j]-97]++
	}

	mapAna[alph] = append(mapAna[alph], strs[i])
 }

 var result [][]string

 for _, v := range mapAna{
	result = append(result, v)
 }

 return result
}
