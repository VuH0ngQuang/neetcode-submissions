func checkInclusion(s1 string, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	n := len(r1)
	if n > len(r2) {
		return false
	}

	need := make(map[rune]int)
	for _, ch := range r1 {
		need[ch]++
	}

	win := make(map[rune]int)
	for i, ch := range r2 {
		win[ch]++ // enter

		if i >= n { // window is now too big, drop the oldest
			old := r2[i-n]
			win[old]--
			if win[old] == 0 {
				delete(win, old) // required, or map equality fails
			}
		}

		if i >= n-1 && equal(need, win) {
			return true
		}
	}
	return false
}

func equal(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}