func maxArea(heights []int) int {
	var ans, left, right = 0, 0, len(heights) - 1
	var templ, tempr int
	for left < right {
		templ = heights[left]
		tempr = heights[right]
		ans = max(ans, min(templ,tempr) * (right - left))
		if templ > tempr {
			right --
		} else {
			left++
		}
	}

	return ans
}
