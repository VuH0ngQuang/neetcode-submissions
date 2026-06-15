func characterReplacement(s string, k int) int {
	mp := make(map[rune]int)
	str := []rune(s)
	left, ans, maxFreq := 0, 0, 0

	for right :=0; right < len(str); right++ {
		mp[str[right]]++
		maxFreq = max(maxFreq, mp[str[right]])

		windowSize := right - left + 1
		need := windowSize - maxFreq

		if need > k {
			mp[str[left]]--
			left++

			windowSize = right - left + 1
			need = windowSize - maxFreq
		}

		ans = max (ans, right - left +1)
	}

	return ans
}
