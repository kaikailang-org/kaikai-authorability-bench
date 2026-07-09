func Rotate(xs []int64, k int) []int64 {
	if len(xs) == 0 { return []int64{} }
	r := k % len(xs)
	out := append([]int64{}, xs[r:]...)
	return append(out, xs[:r]...)
}
