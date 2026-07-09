import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	stock := []Stock{{1, 5}, {2, 0}, {3, 12}, {2, 1}}
	ok := eqI64(InventoryLow(stock, 5), []int64{2, 2}) &&
		eqI64(InventoryLow(stock, 6), []int64{1, 2, 2}) &&
		eqI64(InventoryLow(stock, 0), []int64{}) &&
		eqI64(InventoryLow([]Stock{}, 10), []int64{})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
