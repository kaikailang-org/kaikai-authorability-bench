import "fmt"
import "os"

func itEq(a, b Item) bool { return a == b }

func main() {
	ok := true

	inv, o := AddStock([]Item{}, 7, 10)
	ok = ok && o && len(inv) == 1 && itEq(inv[0], Item{7, 10, 0})

	inv, o = Reserve(inv, 7, 4)
	ok = ok && o && itEq(inv[0], Item{7, 10, 4})

	inv, o = Ship(inv, 7, 3)
	ok = ok && o && itEq(inv[0], Item{7, 7, 1})

	av, o := Available(inv, 7)
	ok = ok && o && av == 6

	// order preserved, append at end
	inv2, _ := AddStock(inv, 3, 5)
	inv2, _ = AddStock(inv2, 7, 2)
	ok = ok && len(inv2) == 2 && inv2[0].Sku == 7 && inv2[0].Qty == 9 && inv2[1].Sku == 3

	// purity: reserve failure leaves input untouched
	before := inv2[0]
	_, bad := Reserve(inv2, 7, 100)
	ok = ok && !bad && itEq(inv2[0], before)

	// validations
	_, v1 := AddStock(inv2, 1, 0)
	_, v2 := Reserve(inv2, 99, 1)
	_, v3 := Reserve(inv2, 7, 9) // available is 8
	_, v4 := Release(inv2, 7, 2) // reserved is 1
	_, v5 := Ship(inv2, 7, 2)    // reserved is 1
	_, v6 := Available(inv2, 42)
	ok = ok && !v1 && !v2 && !v3 && !v4 && !v5 && !v6

	// reserve exactly available, ship to zero, item stays
	inv3, _ := AddStock([]Item{}, 1, 2)
	inv3, o = Reserve(inv3, 1, 2)
	ok = ok && o
	inv3, o = Ship(inv3, 1, 2)
	ok = ok && o && len(inv3) == 1 && itEq(inv3[0], Item{1, 0, 0})

	// release undoes a hold
	inv4, _ := AddStock([]Item{}, 2, 3)
	inv4, _ = Reserve(inv4, 2, 3)
	inv4, o = Release(inv4, 2, 3)
	av4, _ := Available(inv4, 2)
	ok = ok && o && av4 == 3

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
