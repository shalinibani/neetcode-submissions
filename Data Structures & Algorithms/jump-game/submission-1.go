func canJump(nums []int) bool {
	furthest := 0

	for i := 0; i<len(nums); i++{
		if i > furthest{
			return false
		}

		furthest = max(furthest, nums[i] + i)

		if furthest >= len(nums)-1{
			return true
		}
	}

	return true
}
