func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix[0]) - 1

	for _,arr := range matrix {
		if arr[row] == target || arr[0] == target {
			return true
		}
		if arr[row] > target {
			return search(arr, target)
		}
	}

	return false
}

func search(arr []int, target int) bool {
	left, mid, right := 0, 0, len(arr) - 1

	for left <= right {
		mid = ((right - left)/2) + left

		if arr[mid] == target {
			return true
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
