type car struct{
	pos int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]car, 0, len(position))

	for i, p := range position{
		cars = append(cars, car{p, speed[i]})
	}

	sort.Slice(cars, func(i,j int)bool{
		return cars[i].pos > cars[j].pos
	})

	var stack []float64

	for _, c := range cars{
		time := float64(target - c.pos)/float64(c.speed )

		if len(stack) > 0 && stack[len(stack)-1] >= time{
			continue
		}

		stack = append(stack, time)
	} 

	return len(stack)
}
