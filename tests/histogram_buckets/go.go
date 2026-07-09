import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	ok := eqI64(HistogramBuckets([]int64{5,15,25,99,0,9}), []int64{3,1,1,0,0,0,0,0,0,1}) &&
		eqI64(HistogramBuckets([]int64{}), make([]int64, 10)) &&
		eqI64(HistogramBuckets([]int64{10,19,90}), []int64{0,2,0,0,0,0,0,0,0,1}) &&
		eqI64(HistogramBuckets([]int64{0}), []int64{1,0,0,0,0,0,0,0,0,0})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
