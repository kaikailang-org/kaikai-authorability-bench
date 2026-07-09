func TitleCase(s string) string {
	out := []rune{}
	inWord := false
	for _, c := range s {
		if c == ' ' {
			inWord = false
			out = append(out, c)
			continue
		}
		if inWord {
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
		} else {
			if c >= 'a' && c <= 'z' {
				c -= 'a' - 'A'
			}
		}
		inWord = true
		out = append(out, c)
	}
	return string(out)
}
