import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	ok := eqI64(Rotate([]int64{1,2,3,4}, 1), []int64{2,3,4,1}) &&
		eqI64(Rotate([]int64{1,2,3,4}, 0), []int64{1,2,3,4}) &&
		eqI64(Rotate([]int64{1,2,3,4}, 4), []int64{1,2,3,4}) &&
		eqI64(Rotate([]int64{1,2,3,4}, 6), []int64{3,4,1,2}) &&
		eqI64(Rotate([]int64{}, 3), []int64{})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
