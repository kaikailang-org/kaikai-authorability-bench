func SplitBill(total, n int64) []int64 {
	base, r := total/n, total%n
	out := make([]int64, n)
	for i := int64(0); i < n; i++ {
		out[i] = base
		if i < r {
			out[i]++
		}
	}
	return out
}
