func NextFreeID(taken []int64) int64 {
	for c := int64(0); ; c++ {
		used := false
		for _, t := range taken {
			if t == c {
				used = true
				break
			}
		}
		if !used {
			return c
		}
	}
}
