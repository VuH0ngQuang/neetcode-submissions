func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))
	
	for _, str := range tokens {
		if str == "+" || str == "-" || str == "*" || str == "/" {
			b := nums[len(nums) - 1]
			a := nums[len(nums) - 2]
			nums = nums[:len(nums) - 2]
			var val int
			switch str {
				case "+":
					val = a + b

				case "-":
					val = a - b

				case "*":
					val = a * b

				case "/":
					val = a / b
			}
			nums = append(nums, val)
		} else {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
	}

	return nums[len(nums) - 1]
}
