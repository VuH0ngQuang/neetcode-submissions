func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1;

	for left != right {
		tmp := numbers[left] + numbers[right]

		if tmp > target {
			right--
		}
		if tmp < target {
			left++
		}
		if tmp == target {
			return []int{left+1, right+1}
		}
	}

	return []int{0,0}
}
