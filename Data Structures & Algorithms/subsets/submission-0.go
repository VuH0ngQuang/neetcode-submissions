func subsets(nums []int) [][]int {
	result := [][]int{}
	current := []int{}

	var dfs func (int)
	dfs = func (start int) {
		cpy := make([]int, len(current))
		copy(cpy, current)
		result = append(result, cpy)

		for i:= start; i < len (nums); i++ {
			current = append(current, nums[i])
			dfs(i+1)
			current = current[:len(current) - 1]
		}
	}

	dfs(0)
	return result
}
