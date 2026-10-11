func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, res := prices[0], 0

	for _, val := range prices[1:] {
		if val < minPrice {
			minPrice = val
		} else {
			res = max(res, val-minPrice)
		}
	}

	return res
}