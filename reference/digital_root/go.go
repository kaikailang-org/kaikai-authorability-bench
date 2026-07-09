func DigitalRoot(n int64) int64 {
	for n >= 10 {
		var s int64
		for m := n; m > 0; m /= 10 {
			s += m % 10
		}
		n = s
	}
	return n
}
