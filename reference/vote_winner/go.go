func VoteWinner(votes []int64) (int64, bool) {
	var winner int64
	best := 0
	tied := false
	seen := []int64{}
	for _, c := range votes {
		dup := false
		for _, s := range seen {
			if s == c {
				dup = true
				break
			}
		}
		if dup {
			continue
		}
		seen = append(seen, c)
		n := 0
		for _, x := range votes {
			if x == c {
				n++
			}
		}
		if n > best {
			best, winner, tied = n, c, false
		} else if n == best {
			tied = true
		}
	}
	if tied || best == 0 {
		return 0, false
	}
	return winner, true
}
