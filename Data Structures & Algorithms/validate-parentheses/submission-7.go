func isValid(s string) bool {
    var stack []rune

	for _, ch := range s {
		if ch == '(' ||
			ch == '{' ||
			ch == '[' {
				stack = append(stack, ch)
		}

		if len(stack) == 0 {
			return false 
		}

		if ch == ')' || ch == '}' || ch == ']' {
				if ch == ')' && stack[len(stack) - 1] == '(' {
					stack = stack[:len(stack)-1]
					continue
				}
				if ch == '}' && stack[len(stack) - 1] == '{' {
					stack = stack[:len(stack)-1]
					continue
				}
				if ch == ']' && stack[len(stack) - 1] == '[' {
					stack = stack[:len(stack)-1]
					continue
				}
				return false
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
