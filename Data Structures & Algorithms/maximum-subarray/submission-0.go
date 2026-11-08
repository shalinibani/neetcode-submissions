func maxSubArray(nums []int) int {
	maxSum,sum := nums[0], nums[0]

	for i:= 1 ; i<len(nums); i++{
		sum = max(nums[i], sum+nums[i])
		maxSum = max(sum, maxSum)
	}

	return maxSum
}
