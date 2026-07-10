type Account struct{ Id, Balance int64 }

func findAcc(ledger []Account, id int64) int {
	for i, a := range ledger {
		if a.Id == id {
			return i
		}
	}
	return -1
}

func cloneAcc(ledger []Account) []Account {
	return append([]Account{}, ledger...)
}

func OpenAccount(ledger []Account, id int64) ([]Account, bool) {
	if findAcc(ledger, id) >= 0 {
		return nil, false
	}
	return append(cloneAcc(ledger), Account{id, 0}), true
}

func Deposit(ledger []Account, id, n int64) ([]Account, bool) {
	i := findAcc(ledger, id)
	if n < 1 || i < 0 {
		return nil, false
	}
	out := cloneAcc(ledger)
	out[i].Balance += n
	return out, true
}

func Withdraw(ledger []Account, id, n int64) ([]Account, bool) {
	i := findAcc(ledger, id)
	if n < 1 || i < 0 || ledger[i].Balance < n {
		return nil, false
	}
	out := cloneAcc(ledger)
	out[i].Balance -= n
	return out, true
}

func Transfer(ledger []Account, from, to, n int64) ([]Account, bool) {
	fi, ti := findAcc(ledger, from), findAcc(ledger, to)
	if n < 1 || fi < 0 || ti < 0 || from == to || ledger[fi].Balance < n {
		return nil, false
	}
	out := cloneAcc(ledger)
	out[fi].Balance -= n
	out[ti].Balance += n
	return out, true
}

func BalanceOf(ledger []Account, id int64) (int64, bool) {
	i := findAcc(ledger, id)
	if i < 0 {
		return 0, false
	}
	return ledger[i].Balance, true
}
