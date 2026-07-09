func CountPeaks(xs []int64) int64 {
	var n int64
	for i := 1; i+1 < len(xs); i++ {
		if xs[i] > xs[i-1] && xs[i] > xs[i+1] {
			n++
		}
	}
	return n
}
