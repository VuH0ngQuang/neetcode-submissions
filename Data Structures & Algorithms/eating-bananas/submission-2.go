import "slices"

func minEatingSpeed(piles []int, h int) int {
	slices.Sort(piles)
	var left, right = 1, slices.Max(piles)
	for left <= right {
		var tmp int = 0;
		var mid = left + (right - left)/2
		for _,val := range piles {
			tmp += val/mid
			if val % mid != 0 {
				tmp++
			}
		}
		if tmp <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}	
	return left
}
