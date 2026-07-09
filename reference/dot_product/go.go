func DotProduct(a, b []int64) (int64, bool) {
	if len(a) != len(b) { return 0, false }
	var s int64
	for i := range a { s += a[i] * b[i] }
	return s, true
}
