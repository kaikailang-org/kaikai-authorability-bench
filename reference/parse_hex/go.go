func ParseHex(s string) (int64, bool) {
	if len(s) < 3 || s[0] != '0' || s[1] != 'x' {
		return 0, false
	}
	var acc int64
	for _, c := range s[2:] {
		var d int64
		switch {
		case c >= '0' && c <= '9':
			d = int64(c - '0')
		case c >= 'a' && c <= 'f':
			d = int64(c-'a') + 10
		case c >= 'A' && c <= 'F':
			d = int64(c-'A') + 10
		default:
			return 0, false
		}
		acc = acc*16 + d
	}
	return acc, true
}
