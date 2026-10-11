func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	last := make(map[rune]int)
	left, ans := 0, 0

	for right, ch := range runes {
		if idx, exist := last[ch]; exist && idx >= left {
			left = idx + 1
		}
		last[ch] = right
		if size := right - left + 1; size > ans {
			ans = size
		}
	}
	return ans
}