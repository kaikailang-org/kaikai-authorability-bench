func CollatzSteps(n int64) int64 {
	var s int64
	for n != 1 {
		if n%2 == 0 { n = n / 2 } else { n = 3*n + 1 }
		s++
	}
	return s
}
