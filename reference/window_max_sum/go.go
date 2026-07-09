func WindowMaxSum(xs []int64, k int) (int64, bool) {
	if k <= 0 || len(xs) < k {
		return 0, false
	}
	best := int64(0)
	first := true
	for i := 0; i+k <= len(xs); i++ {
		var sum int64
		for j := i; j < i+k; j++ {
			sum += xs[j]
		}
		if first || sum > best {
			best = sum
			first = false
		}
	}
	return best, true
}
