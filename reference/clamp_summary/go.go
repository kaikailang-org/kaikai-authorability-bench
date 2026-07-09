type Clamped struct {
	Values  []int64
	Changed int64
}

func ClampSummary(xs []int64, lo, hi int64) Clamped {
	values := []int64{}
	var changed int64
	for _, x := range xs {
		v := x
		if v < lo {
			v = lo
		}
		if v > hi {
			v = hi
		}
		if v != x {
			changed++
		}
		values = append(values, v)
	}
	return Clamped{Values: values, Changed: changed}
}
