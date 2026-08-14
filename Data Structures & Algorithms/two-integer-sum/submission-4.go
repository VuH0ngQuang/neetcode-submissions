func twoSum(nums []int, target int) []int {
    var hmap = make(map[int]int, len(nums))

	for i,v := range nums {
		var need =  target - v;
		if j, ok := hmap[need]; ok {
			return []int{j, i}
		}
		hmap[v] = i;
	}
	
	return []int{0,0}
}
