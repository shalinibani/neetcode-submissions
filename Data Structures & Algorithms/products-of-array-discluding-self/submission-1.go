func productExceptSelf(nums []int) []int {
prefix := make([]int, len(nums))
suffix := make([]int, len(nums))

for i:=0; i<len(prefix); i++{
	prefix[i] = 1
	suffix[i] = 1
}


for i := 1; i<len(nums); i++{
 prefix[i] = prefix[i-1]*nums[i-1]
}

for i := len(nums)-2;i>=0; i--{
	suffix[i] = suffix[i+1] * nums[i+1]
}

res := make([]int, len(nums))

for i:= 0; i<len(nums) ; i++{
 res[i] = prefix[i] * suffix[i]
}

 return res
}
