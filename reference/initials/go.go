func Initials(name string) string {
	out := []rune{}
	inWord := false
	for _, c := range name {
		if c == ' ' {
			inWord = false
			continue
		}
		if !inWord {
			if c >= 'a' && c <= 'z' {
				c -= 'a' - 'A'
			}
			out = append(out, c)
		}
		inWord = true
	}
	return string(out)
}
