import "sort"

func MedianFloor(xs []int64) (int64, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	s := append([]int64{}, xs...)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	n := len(s)
	if n%2 == 1 {
		return s[n/2], true
	}
	return (s[n/2-1] + s[n/2]) / 2, true
}
