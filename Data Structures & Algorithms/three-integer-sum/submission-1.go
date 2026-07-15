import "slices"
func threeSum(nums []int) [][]int {
var res [][]int

 for i:=0; i<=len(nums)-3; i++{
	twoSumSol := twoSum(nums[i+1:], nums[i]*-1)
	if len(twoSumSol) > 0{
		for _, sol := range twoSumSol{
			res = append(res, sol)
		}
		
	}
	
 }

 resMap := make(map[[3]int]bool)
 var finalRes[][]int

for _, sol := range res{
	key := [3]int{sol[0], sol[1], sol[2]}

	slices.Sort(key[:])

	if !resMap[key]{
		resMap[key] = true
		finalRes = append(finalRes, sol)
	}
}

 return finalRes
}

func twoSum(nums[]int, target int)[][]int{
	hashMap := make(map[int]struct{})

	var res  [][]int

	for _, n := range nums{
		rem := target - n

		if _, ok := hashMap[rem]; ok{
				res = append(res, []int{target*-1, n, rem})
			
		}

		hashMap[n] = struct{}{}
	}

	return res
}
