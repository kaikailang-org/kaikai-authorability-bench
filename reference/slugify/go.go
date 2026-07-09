func Slugify(s string) string {
	out := []rune{}
	pending := false
	for _, c := range s {
		isAlnum := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
		if isAlnum {
			if pending && len(out) > 0 {
				out = append(out, '-')
			}
			pending = false
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
			out = append(out, c)
		} else {
			pending = true
		}
	}
	return string(out)
}
