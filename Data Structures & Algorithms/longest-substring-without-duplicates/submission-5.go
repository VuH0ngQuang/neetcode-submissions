func lengthOfLongestSubstring(s string) int {
	left, ans := 0, 0
	seen := map[byte]int{}

	for right := 0; right < len(s); right++ {
		if prev, ok := seen[s[right]]; ok && prev >= left {
			left = prev + 1
		}

		seen[s[right]] = right
		ans = max(ans, right-left+1)
	}

	return ans
}