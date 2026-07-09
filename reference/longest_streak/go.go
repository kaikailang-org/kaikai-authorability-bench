func LongestStreak(days []bool) int64 {
	var cur, best int64
	for _, d := range days {
		if d { cur++ } else { cur = 0 }
		if cur > best { best = cur }
	}
	return best
}
