import "slices"

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, slices.Max(piles)
	rate := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2

		hours := 0
		for _, pile := range piles {
			hours += (pile + mid - 1) / mid
		}

		if hours <= h {
			rate = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return rate
}