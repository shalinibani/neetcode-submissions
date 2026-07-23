import "slices"
func minEatingSpeed(piles []int, h int) int {
   left, right := 1, slices.Max(piles)
   for left <= right{
	mid := left + (right -left)/2
	currh := 0
	for _, p := range piles{
		if currh > h{
			break
		}

		currh += (p+mid-1)/mid
	}

	if currh <= h{
		right = mid - 1
	}else{
		left = mid + 1
	}
   }

   return left
}
