func twoSum(nums []int, target int) []int {
    var hmap = make(map[int]int)

	for i,v := range nums {
		var need =  target - v;
		if _, ok := hmap[need]; ok {
			return []int{hmap[need], i}
		}
		hmap[v] = i;
	}
	
	return []int{0,0}
}
