func isPalindrome(s string) bool {
	var str []rune = []rune(s)
	var left, right int = 0, len(str)-1
	for left < right {
		if !unicode.IsLetter(str[left]) && !unicode.IsDigit(str[left]) {
			left++
			continue
		}
		if !unicode.IsLetter(str[right]) && !unicode.IsDigit(str[right]) {
			right--
			continue
		}

		if unicode.ToLower(str[left]) != unicode.ToLower(str[right]) {
			return false
		}
		left++
		right--
	}

	return true
}
