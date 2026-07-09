import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	ok := eqI64(Interleave([]int64{1,3,5}, []int64{2,4}), []int64{1,2,3,4,5}) &&
		eqI64(Interleave([]int64{}, []int64{7,8}), []int64{7,8}) &&
		eqI64(Interleave([]int64{1}, []int64{}), []int64{1}) &&
		eqI64(Interleave([]int64{}, []int64{}), []int64{}) &&
		eqI64(Interleave([]int64{1}, []int64{7,8,9}), []int64{1,7,8,9})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
