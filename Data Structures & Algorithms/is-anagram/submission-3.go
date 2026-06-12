func isAnagram(s string, t string) bool {

	if (len(s) != len(t)) {
		return false
	}

	var m = make(map[rune]int)

	for _,char :=  range s {
		m[char] += 1
	}

	for _,char := range t {
		value, exist := m[char]
		if exist {
			if value == 0 {
				return false
			} else {
				m[char] -=1
			}
		} else {
			return false
		}
	}

	return true
}
