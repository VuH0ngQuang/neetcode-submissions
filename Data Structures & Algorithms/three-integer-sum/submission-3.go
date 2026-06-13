import "slices"

func threeSum(nums []int) [][]int {
	var ans [][]int
	var left, right, temp int

	slices.Sort(nums)

	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left = i + 1;
		right = len(nums) - 1;
		for ( left < right) {
			temp = nums[i]+nums[left]+nums[right]
			if temp > 0 {
				right --
			} else if temp < 0 {
				left ++
			} else if temp == 0 {
				ans = append(ans,[]int{nums[i],nums[left],nums[right]})
				left ++
				right --
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return ans
}
