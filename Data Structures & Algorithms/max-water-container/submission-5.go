func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1
	res, tmp := 0, 0
	
	for left < right {
		tmp = min(heights[left], heights[right]) * (right - left)
		res = max(res, tmp)
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

