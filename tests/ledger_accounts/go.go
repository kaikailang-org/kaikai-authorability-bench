import "fmt"
import "os"

func bal(l []Account, id, want int64) bool {
	v, o := BalanceOf(l, id)
	return o && v == want
}

func main() {
	ok := true

	l, o := OpenAccount([]Account{}, 1)
	ok = ok && o
	l, o = OpenAccount(l, 2)
	ok = ok && o && len(l) == 2 && l[0].Id == 1 && l[1].Id == 2

	_, dup := OpenAccount(l, 1)
	ok = ok && !dup

	l, o = Deposit(l, 1, 100)
	ok = ok && o && bal(l, 1, 100)

	l, o = Transfer(l, 1, 2, 40)
	ok = ok && o && bal(l, 1, 60) && bal(l, 2, 40)

	_, b1 := Transfer(l, 2, 2, 10)
	_, b2 := Transfer(l, 1, 9, 10)
	_, b3 := Transfer(l, 9, 1, 10)
	_, b4 := Transfer(l, 2, 1, 41)
	_, b5 := Transfer(l, 1, 2, 0)
	ok = ok && !b1 && !b2 && !b3 && !b4 && !b5

	l2, o2 := Withdraw(l, 2, 40)
	ok = ok && o2 && bal(l2, 2, 0)
	_, b6 := Withdraw(l2, 2, 1)
	_, b7 := Withdraw(l2, 1, 0)
	_, b8 := Withdraw(l2, 3, 5)
	ok = ok && !b6 && !b7 && !b8

	_, b9 := Deposit(l2, 1, -5)
	_, b10 := Deposit(l2, 42, 5)
	ok = ok && !b9 && !b10

	// purity: failed transfer leaves ledger intact
	ok = ok && bal(l, 1, 60) && bal(l, 2, 40)

	_, b11 := BalanceOf(l, 99)
	ok = ok && !b11

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
