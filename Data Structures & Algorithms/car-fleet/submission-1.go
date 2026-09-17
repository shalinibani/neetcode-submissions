type car struct{
	pos int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	s := make([]car, 0, len(position))

	for i, p := range position {
		s = append(s, car{p, speed[i]})
	}

	sort.Slice(s, func(i,j int)bool{
		return s[i].pos > s[j].pos
	})

	var s1 []float64 

	for _, car := range s{
		time := (float64(target) - float64(car.pos))/float64(car.speed)

		if len(s1) > 0 && s1[len(s1)-1] >= time{
			continue

		}else{
			s1 = append(s1, time)
		}
	}

	return len(s1)

}
