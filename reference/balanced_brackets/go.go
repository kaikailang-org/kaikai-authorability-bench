func BalancedBrackets(s string) bool {
	stack := []rune{}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			want := '('
			if c == ']' {
				want = '['
			} else if c == '}' {
				want = '{'
			}
			if len(stack) == 0 || stack[len(stack)-1] != want {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
