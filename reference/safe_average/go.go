func SafeAverage(xs []int64) (int64, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	var sum int64
	for _, x := range xs {
		sum += x
	}
	return sum / int64(len(xs)), true
}
