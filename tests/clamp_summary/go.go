import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	a := ClampSummary([]int64{1, 50, -3, 7}, 0, 10)
	b := ClampSummary([]int64{}, 0, 10)
	c := ClampSummary([]int64{5, 5}, 5, 5)
	ok := eqI64(a.Values, []int64{1, 10, 0, 7}) && a.Changed == 2 &&
		eqI64(b.Values, []int64{}) && b.Changed == 0 &&
		eqI64(c.Values, []int64{5, 5}) && c.Changed == 0
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
