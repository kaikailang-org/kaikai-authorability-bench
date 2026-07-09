import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	pass := true
	if v, ok := WindowMaxSum([]int64{1,2,3,4}, 2); !chk(v, ok, 7, true) { pass = false }
	if v, ok := WindowMaxSum([]int64{2,-1,2,3}, 2); !chk(v, ok, 5, true) { pass = false }
	if v, ok := WindowMaxSum([]int64{1,2}, 3); !chk(v, ok, 0, false) { pass = false }
	if v, ok := WindowMaxSum([]int64{5}, 1); !chk(v, ok, 5, true) { pass = false }
	if v, ok := WindowMaxSum([]int64{-1,-2,-3}, 1); !chk(v, ok, -1, true) { pass = false }
	if v, ok := WindowMaxSum([]int64{4,4,4}, 3); !chk(v, ok, 12, true) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
