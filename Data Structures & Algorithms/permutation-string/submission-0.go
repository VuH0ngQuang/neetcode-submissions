func checkInclusion(s1 string, s2 string) bool {
	leng := len([]rune(s1))
	mp := make(map[rune]int) 
	str := []rune(s2)
	left := 0

	for _,ch := range s1 {
		mp[ch]++
	}

	for right :=0; right < len(str); right++ {
		mp[str[right]]--
		if mp[str[right]] == 0 {
			delete(mp, str[right])
		}
		windowSize := right - left  + 1
		
		if windowSize > leng {
			mp[str[left]]++
			if mp[str[left]] == 0 {
				delete(mp, str[left])
			}
			left++
		}

		if len(mp) == 0 {
			return true
		}
	}

	return false
}
