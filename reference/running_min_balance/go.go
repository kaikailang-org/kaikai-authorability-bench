func RunningMinBalance(txs []int64) int64 {
	var bal, low int64
	for _, t := range txs {
		bal += t
		if bal < low {
			low = bal
		}
	}
	return low
}
