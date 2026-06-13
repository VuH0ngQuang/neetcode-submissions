func evalRPN(tokens []string) int {
	var stack = make([]int,0, len(tokens))
	var tmp1, tmp2 int
	for _, val := range tokens {
		switch val {
			case "+":
				tmp2 = stack[len(stack) - 1]
				tmp1 = stack[len(stack) - 2]
				stack = stack[:len(stack) - 2]
				stack = append(stack,tmp1+tmp2)
			case "-":
				tmp2 = stack[len(stack) - 1]
				tmp1 = stack[len(stack) - 2]
				stack = stack[:len(stack) - 2]
				stack = append(stack,tmp1-tmp2)
			case "*":
				tmp2 = stack[len(stack) - 1]
				tmp1 = stack[len(stack) - 2]
				stack = stack[:len(stack) - 2]
				stack = append(stack,tmp1*tmp2)
			case "/":
				tmp2 = stack[len(stack) - 1]
				tmp1 = stack[len(stack) - 2]
				stack = stack[:len(stack) - 2]
				stack = append(stack,tmp1/tmp2)
			default:
				num, _ := strconv.Atoi(val)
				stack = append(stack,num)
		}
	}
	return stack[0]
}
