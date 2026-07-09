func Caesar(s string, k int64) string {
	r := rune(((k % 26) + 26) % 26)
	out := []rune{}
	for _, c := range s {
		switch {
		case c >= 'a' && c <= 'z':
			out = append(out, 'a'+(c-'a'+r)%26)
		case c >= 'A' && c <= 'Z':
			out = append(out, 'A'+(c-'A'+r)%26)
		default:
			out = append(out, c)
		}
	}
	return string(out)
}
