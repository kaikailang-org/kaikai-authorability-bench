func RangeOverlap(a, b Range) (Range, bool) {
	lo := a.Start
	if b.Start > lo {
		lo = b.Start
	}
	hi := a.Finish
	if b.Finish < hi {
		hi = b.Finish
	}
	if lo < hi {
		return Range{Start: lo, Finish: hi}, true
	}
	return Range{}, false
}

type Range struct{ Start, Finish int64 }
