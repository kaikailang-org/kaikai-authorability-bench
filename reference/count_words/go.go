func CountWords(s string) int64 {
	var n int64
	inWord := false
	for _, c := range s {
		sp := c == ' ' || c == '\t'
		if !sp && !inWord {
			n++
		}
		inWord = !sp
	}
	return n
}
