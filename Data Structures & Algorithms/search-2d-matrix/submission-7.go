func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1

	for l <= r {
		m := l + (r-l)/2
		v := matrix[m/cols][m%cols]
		if v == target {
			return true
		}
		if v < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}