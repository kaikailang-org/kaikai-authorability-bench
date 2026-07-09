func HistogramBuckets(xs []int64) []int64 {
	out := make([]int64, 10)
	for _, x := range xs {
		out[x/10]++
	}
	return out
}
