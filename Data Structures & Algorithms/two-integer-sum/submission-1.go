func twoSum(nums []int, target int) []int {
    var mp = make(map[int]int)
	var temp int
	mp[nums[0]] = 0

	for i := 1; i < len(nums); i++ {
		temp = target - nums[i]
		index, exist := mp[temp]
		if exist {
			return []int{index, i}
		} else {
			mp[nums[i]] = i
		}
	}

	return []int{0, 0}
}
