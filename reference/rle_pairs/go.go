type Run struct{ V, N int64 }

func Rle(xs []int64) []Run {
	out := []Run{}
	for _, x := range xs {
		if len(out) > 0 && out[len(out)-1].V == x {
			out[len(out)-1].N++
		} else {
			out = append(out, Run{V: x, N: 1})
		}
	}
	return out
}
