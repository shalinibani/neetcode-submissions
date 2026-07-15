func maxProfit(prices []int) int {
 maxProfi := 0

 if len(prices) <= 1{
	return maxProfi
 }

 left, right := 0,0

	for right < len(prices){
		if prices[right] > prices[left]{
			maxProfi = max(maxProfi, prices[right] - prices[left])
			
		}else{
			left = right
		}

		right++

	}

	return maxProfi

}

func max(a, b int) int{
	if a > b{
		return a
	}

	return b
}
