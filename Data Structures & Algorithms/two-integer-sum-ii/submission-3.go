func twoSum(numbers []int, target int) []int {
	var mp = make(map[int]int)
	var temp int 
	for index, number := range numbers {
		temp = target - number
		idx, exist := mp[temp]
		if exist {
			return []int{idx + 1, index + 1}
		} else {
			mp[number] = index
		}
	}

	return []int{0,0}
}
