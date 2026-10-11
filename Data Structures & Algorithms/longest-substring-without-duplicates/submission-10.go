func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	mp := make(map[rune]int)
	left, res := 0, 0

	for right, ch := range runes {
		if idx, exists := mp[ch]; exists && idx >= left {
			left = idx + 1
		}
		mp[ch] = right
		res = max (res, right - left + 1)
	}

	return res
}
