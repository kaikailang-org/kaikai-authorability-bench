func MaskDigits(s string) string {
	nd := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			nd++
		}
	}
	toMask := 0
	if nd > 4 {
		toMask = nd - 4
	}
	out := []rune{}
	for _, c := range s {
		if c >= '0' && c <= '9' && toMask > 0 {
			out = append(out, '*')
			toMask--
		} else {
			out = append(out, c)
		}
	}
	return string(out)
}
