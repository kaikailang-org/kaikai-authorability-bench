import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	ok := eqI64(RleDecode([]Run{{5,3},{2,1}}), []int64{5,5,5,2}) &&
		eqI64(RleDecode([]Run{}), []int64{}) &&
		eqI64(RleDecode([]Run{{7,0},{1,2}}), []int64{1,1}) &&
		eqI64(RleDecode([]Run{{-3,2}}), []int64{-3,-3})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
