func pop(stack []int) (int, []int) {
	return stack[len(stack) - 1], stack[:len(stack)-1]
}

func evalRPN(tokens []string) int {
	var stack = make([]int,0)
	var tmp1, tmp2 int
	for _, val := range tokens {
		switch val {
			case "+":
				tmp2, stack = pop(stack)
				tmp1, stack = pop(stack)
				stack = append(stack,tmp1+tmp2)
			case "-":
				tmp2, stack = pop(stack)
				tmp1, stack = pop(stack)
				stack = append(stack,tmp1-tmp2)
			case "*":
				tmp2, stack = pop(stack)
				tmp1, stack = pop(stack)
				stack = append(stack,tmp1*tmp2)
			case "/":
				tmp2, stack = pop(stack)
				tmp1, stack = pop(stack)
				stack = append(stack,tmp1/tmp2)
			default:
				num, _ := strconv.Atoi(val)
				stack = append(stack,num)
		}
	}
	return stack[0]
}
