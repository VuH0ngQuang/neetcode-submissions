func maxProfit(prices []int) int {
	ans := 0
	left := 0

	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			left = right
		} else {
			ans = max(ans, prices[right]-prices[left])
		}
	}

	return ans
}