type stack struct{
	temp int
	idx int
}

func dailyTemperatures(temperatures []int) []int {
  var s []stack
  res := make([]int, len(temperatures))

  for i, t := range temperatures {
	for len(s) > 0 && s[len(s)-1].temp < t{
		n := s[len(s)-1]
		s = s[:len(s)-1]

		res[n.idx] = i - n.idx
	}

	s = append(s, stack{t, i})
  }

return res
}
