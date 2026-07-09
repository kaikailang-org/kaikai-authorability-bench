import "sort"

type Range struct{ Start, Finish int64 }

func MergeRanges(rs []Range) []Range {
	s := append([]Range{}, rs...)
	sort.Slice(s, func(i, j int) bool { return s[i].Start < s[j].Start })
	out := []Range{}
	for _, r := range s {
		if n := len(out); n > 0 && r.Start <= out[n-1].Finish {
			if r.Finish > out[n-1].Finish {
				out[n-1].Finish = r.Finish
			}
		} else {
			out = append(out, r)
		}
	}
	return out
}
