func maxProfit(prices []int) int {
	maxProf := 0

	left := 0

	for right := 0; right<len(prices) ; right++{
		if prices[left] < prices[right]{
			maxProf = max(maxProf, prices[right]-prices[left])
		} else {
			left = right
		}
	}

	return maxProf
}
