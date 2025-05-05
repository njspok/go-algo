package validparentheses

var openBrace = map[uint8]struct{}{
	'(': {},
	'[': {},
	'{': {},
}

var closeBrace = map[uint8]struct{}{
	')': {},
	']': {},
	'}': {},
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	// check start ({[
	if _, ok := closeBrace[s[0]]; ok {
		return false
	}

	stack := stack[uint8]{}

	for _, ch := range s {
		ch := uint8(ch)
		if _, ok := openBrace[ch]; ok {
			stack.push(ch)
			continue
		}

		// todo check not brace symbols

		// closed brace

		if len(stack) == 0 {
			return false
		}

		last := stack.pop()
		switch {
		case last == '(' && ch == ')':
			continue
		case last == '[' && ch == ']':
			continue
		case last == '{' && ch == '}':
			continue
		default:
			return false
		}

	}

	return len(stack) == 0
}
