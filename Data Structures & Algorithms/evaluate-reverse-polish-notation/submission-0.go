func evalRPN(tokens []string) int {
  stack := make([]string, 0, len(tokens))

  for _, t := range tokens{
	if t == "+" || t == "-" || t == "/" || t == "*"{
		n2, _ := strconv.Atoi(stack[len(stack)-1])
		n1, _ := strconv.Atoi(stack[len(stack)-2])

		stack = stack[:len(stack)-2]

		res := 0

		switch t{
			case "+": res = n1+n2
			case "-": res = n1-n2
			case "*": res = n1*n2
			case "/": res = n1/n2
		}

		stack = append(stack, strconv.Itoa(res))
	}else{
		stack = append(stack, t)
	}
  }

  finalRes, _ := strconv.Atoi(stack[0])
  return finalRes
}
