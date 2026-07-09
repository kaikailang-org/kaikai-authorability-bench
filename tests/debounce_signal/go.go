import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	ok := eqI64(Debounce([]int64{4,4,7,7,7,4}), []int64{4,7,4}) &&
		eqI64(Debounce([]int64{}), []int64{}) &&
		eqI64(Debounce([]int64{9}), []int64{9}) &&
		eqI64(Debounce([]int64{2,2,2,2}), []int64{2}) &&
		eqI64(Debounce([]int64{1,2,3}), []int64{1,2,3})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
