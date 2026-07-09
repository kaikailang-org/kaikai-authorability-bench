import "sort"

type Score struct{ ID, Points int64 }

func LeaderboardTop3(scores []Score) []int64 {
	idx := make([]int, len(scores))
	for i := range idx {
		idx[i] = i
	}
	sort.SliceStable(idx, func(a, b int) bool { return scores[idx[a]].Points > scores[idx[b]].Points })
	out := []int64{}
	for i := 0; i < len(idx) && i < 3; i++ {
		out = append(out, scores[idx[i]].ID)
	}
	return out
}
