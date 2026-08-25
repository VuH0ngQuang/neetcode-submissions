func twoSum(nums []int, target int) []int {
	var indices = make(map[int]int)
	for i, num := range nums {
		var trg = target - num
		if j, exist := indices[trg]; exist {
			return []int{j, i}
		}
		indices[num] = i
	}
	return []int{0,0}
}
