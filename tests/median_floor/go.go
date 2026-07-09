import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	pass := true
	if v, ok := MedianFloor([]int64{3,1,2}); !chk(v, ok, 2, true) { pass = false }
	if v, ok := MedianFloor([]int64{4,1,3,2}); !chk(v, ok, 2, true) { pass = false }
	if v, ok := MedianFloor([]int64{}); !chk(v, ok, 0, false) { pass = false }
	if v, ok := MedianFloor([]int64{5}); !chk(v, ok, 5, true) { pass = false }
	if v, ok := MedianFloor([]int64{-3,-4}); !chk(v, ok, -3, true) { pass = false }
	if v, ok := MedianFloor([]int64{1,1,8,8}); !chk(v, ok, 4, true) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
