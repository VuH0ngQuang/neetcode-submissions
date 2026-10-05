func twoSum(nums []int, target int) []int {
	indices := make(map[int]int)

	for index, val := range nums {
		if i, exists := indices[target - val]; exists {
			return []int{i, index}
		}

		indices[val] = index
	}

	return []int{0,0}
}
