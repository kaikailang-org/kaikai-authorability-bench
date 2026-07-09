func TankLevel(cap int64, deltas []int64) int64 {
	var lvl int64
	for _, d := range deltas {
		lvl += d
		if lvl > cap {
			lvl = cap
		}
		if lvl < 0 {
			lvl = 0
		}
	}
	return lvl
}
