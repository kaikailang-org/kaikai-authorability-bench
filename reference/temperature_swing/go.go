func TemperatureSwing(xs []int64) int64 {
	var best int64
	for i := 1; i < len(xs); i++ {
		d := xs[i] - xs[i-1]
		if d < 0 {
			d = -d
		}
		if d > best {
			best = d
		}
	}
	return best
}
