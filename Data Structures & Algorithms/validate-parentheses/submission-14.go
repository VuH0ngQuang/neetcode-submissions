func isValid(s string) bool {
    stack := make([]rune, 0)
	runes := []rune(s)

	if len(s) == 1 {
		return false
	}

	for _,ch := range runes {
		if isOpen(ch) {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			switch ch {
				case ')':
					if pop != '(' {
						return false
					}
				
				case '}':
					if pop != '{' {
						return false
					}

				case ']':
					if pop != '[' {
						return false
					}
			}
		}
	}
	return len(stack) == 0
}

func isOpen(ch rune) bool {
	return ch == '(' || ch == '{' || ch == '['
}

