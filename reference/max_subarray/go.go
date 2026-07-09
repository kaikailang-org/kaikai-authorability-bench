func MaxSubarray(xs []int64) int64 {
	if len(xs) == 0 { return 0 }
	best, cur := xs[0], xs[0]
	for _, x := range xs[1:] {
		if cur+x > x { cur = cur + x } else { cur = x }
		if cur > best { best = cur }
	}
	return best
}
