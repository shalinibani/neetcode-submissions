func maxArea(heights []int) int {
  left, right := 0, len(heights)-1

  maxWater := 0

  for left < right{
	maxWater = max(maxWater, (right-left) * min(heights[left], heights[right]))

	if heights[left] < heights[right]{
		left++
	}else{
		right--
	}
  }

  return maxWater
}

func min(a, b int) int{
	if a > b{
		return b
	}

	return a
}

func max(a, b int) int{
	if a > b{
		return a
	}

	return b
}
