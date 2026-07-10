type Meeting struct{ Start, Finish int64 }

func MinRooms(ms []Meeting) (int64, bool) {
	for _, m := range ms {
		if m.Start >= m.Finish {
			return 0, false
		}
	}
	var best int64 = 0
	for _, m := range ms {
		var c int64 = 0
		for _, x := range ms {
			if x.Start <= m.Start && m.Start < x.Finish {
				c++
			}
		}
		if c > best {
			best = c
		}
	}
	return best, true
}
