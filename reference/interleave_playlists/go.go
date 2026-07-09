func Interleave(a, b []int64) []int64 {
	out := []int64{}
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) { out = append(out, a[i]) }
		if i < len(b) { out = append(out, b[i]) }
	}
	return out
}
