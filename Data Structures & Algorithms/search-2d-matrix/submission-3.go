func searchMatrix(matrix [][]int, target int) bool {
	for _,mtr := range matrix {
		if mtr[len(mtr)-1] >= target {
			left := 0
			right := len(mtr) - 1
			for left <= right {
				mid := left + (right - left)/2
				if mtr[mid] == target {
					return true
				} else if mtr[left] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
 		} else {
			continue
		}
	}
	return false
}
