import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	ok := eqI64(SplitBill(100, 3), []int64{34,33,33}) &&
		eqI64(SplitBill(100, 4), []int64{25,25,25,25}) &&
		eqI64(SplitBill(5, 6), []int64{1,1,1,1,1,0}) &&
		eqI64(SplitBill(0, 2), []int64{0,0}) &&
		eqI64(SplitBill(7, 1), []int64{7})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
