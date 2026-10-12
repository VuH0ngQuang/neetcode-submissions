func characterReplacement(s string, k int) int {
	runes := []rune(s)
	mp := make(map[rune]int)
	left, maxf, res := 0, 0, 0

	for right, ch := range runes {
		mp[ch]++
		maxf = max(maxf, mp[ch])

		for (right - left + 1) - maxf > k {
			mp[runes[left]]--
			left++
		}

		res = max(res, right - left + 1)
	}

	return res
}
