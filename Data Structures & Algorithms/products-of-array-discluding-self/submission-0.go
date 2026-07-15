func productExceptSelf(nums []int) []int {
 prod := 1

 countZero:= 0

 for _, n := range nums{
	if n== 0 {
		countZero++
		continue
	}
	prod *= n
 }

 res := make([]int,0, len(nums))
 for _, n := range nums{
	if n == 0 && countZero==1{
		res = append(res, prod)
		continue
	}

	if countZero>0{
		res = append(res, 0)
	}else{
		res = append(res, prod/n)
	}
    
	
 }

 return res
}
