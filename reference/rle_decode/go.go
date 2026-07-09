type Run struct{ V, N int64 }

func RleDecode(runs []Run) []int64 {
	out := []int64{}
	for _, r := range runs {
		for i := int64(0); i < r.N; i++ {
			out = append(out, r.V)
		}
	}
	return out
}
