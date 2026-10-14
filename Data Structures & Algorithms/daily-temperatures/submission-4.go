type Temp struct {
	temp int
	index int
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]Temp, len(temperatures))

	for i, val := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].temp < val {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top.index] = i - top.index
		}
		stack = append(stack, Temp{temp: val, index: i})
	}
	return res
}
