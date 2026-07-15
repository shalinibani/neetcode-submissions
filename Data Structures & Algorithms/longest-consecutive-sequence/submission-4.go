func longestConsecutive(nums []int) int {
  if len(nums) == 0{
	return 0
  }

  sort.Slice(nums, func(i, j int)bool{
	return nums[i]< nums[j]
  })

  count := 1
  long := 1

  for i := 1; i<len(nums) ; i++{
	if nums [i] - nums[i-1] == 0{
		continue
	}

	if nums [i] - nums[i-1] == 1{
		count++
	}else{
		if count > long{
			long = count
		}
		count = 1
	}
  }

  if count > long {
	long = count
  }

  return long
}

func abs(num int) int{
	if num < 0{
		return num *-1
	}

	return num
}
