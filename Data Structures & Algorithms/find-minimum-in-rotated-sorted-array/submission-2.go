func findMin(nums []int) int {
   res := nums[0]

   left, right := 0, len(nums)-1

   for left <= right{
	mid := left + (right-left)/2

	res = min(res,nums[mid])
	if nums[left] <= nums[right]{
		res = min(res,nums[left])
		break 
	}else if nums[mid] >= nums[left]{
		left = mid + 1
	}else{
		right = mid -1
	}
   }

   return res
}

func min(a, b int) int{
	if a < b{
		return a
	}

	return b
}
