func CleanPalindrome(s string) bool {
	cs := []rune{}
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
			cs = append(cs, c)
		case c >= 'a' && c <= 'z':
			cs = append(cs, c)
		case c >= 'A' && c <= 'Z':
			cs = append(cs, c+('a'-'A'))
		}
	}
	for i, j := 0, len(cs)-1; i < j; i, j = i+1, j-1 {
		if cs[i] != cs[j] {
			return false
		}
	}
	return true
}
