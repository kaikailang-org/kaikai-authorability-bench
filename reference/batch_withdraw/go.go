func BatchWithdraw(balance int64, amounts []int64) (int64, bool) {
	bal := balance
	for _, a := range amounts {
		if a <= 0 || a > bal {
			return 0, false
		}
		bal -= a
	}
	return bal, true
}
