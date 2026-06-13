func evalRPN(tokens []string) int {
	var stack = make([]int, 0, len(tokens))
	for _, val := range tokens {
		switch val {
			case "+":
				n := len(stack)
				tmp2 := stack[n - 1]
				tmp1 := stack[n - 2]
				stack = stack[:n - 2]
				stack = append(stack,tmp1+tmp2)
			case "-":
				n := len(stack)
				tmp2 := stack[n - 1]
				tmp1 := stack[n - 2]
				stack = stack[:n - 2]
				stack = append(stack,tmp1-tmp2)
			case "*":
				n := len(stack)
				tmp2 := stack[n - 1]
				tmp1 := stack[n - 2]
				stack = stack[:n - 2]
				stack = append(stack,tmp1*tmp2)
			case "/":
				n := len(stack)
				tmp2 := stack[n - 1]
				tmp1 := stack[n - 2]
				stack = stack[:n - 2]
				stack = append(stack,tmp1/tmp2)
			default:
				num, _ := strconv.Atoi(val)
				stack = append(stack,num)
		}
	}
	return stack[0]
}
