func dailyTemperatures(temperatures []int) []int {
	var stack []int = make([]int,0,len(temperatures))
	var ans []int = make([]int,len(temperatures))
	
	for i,val := range temperatures {
		for	len(stack) > 0 && val > temperatures[stack[len(stack)-1]] {
			tmp := stack[len(stack)-1]
			ans[tmp] = i - tmp
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ans
}
