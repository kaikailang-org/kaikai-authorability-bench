func AlternatingSum(xs []int64) int64 {
	var sum int64
	sign := int64(1)
	for _, x := range xs {
		sum += sign * x
		sign = -sign
	}
	return sum
}
