func IntSqrt(n int64) int64 {
	lo, hi := int64(0), n
	if hi > 1000000 {
		hi = 1000000
	}
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid*mid <= n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}
