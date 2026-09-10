func evalRPN(tokens []string) int {
  stack := make([]int, 0, len(tokens))

  for _, t := range tokens{
	if t == "+" || t == "-" || t == "/" || t == "*"{
		n2 := stack[len(stack)-1]
		n1 := stack[len(stack)-2]

		stack = stack[:len(stack)-2]

		res := 0

		switch t{
			case "+": res = n1+n2
			case "-": res = n1-n2
			case "*": res = n1*n2
			case "/": res = n1/n2
		}

		stack = append(stack, res)
	}else{
		convT, _ := strconv.Atoi(t)
		stack = append(stack, convT)
	}
  }

  return stack[0]
}
