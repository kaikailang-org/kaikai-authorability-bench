import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	pass := true
	if v, ok := SafeAverage([]int64{1,2,4}); !chk(v, ok, 2, true) { pass = false }
	if v, ok := SafeAverage([]int64{}); !chk(v, ok, 0, false) { pass = false }
	if v, ok := SafeAverage([]int64{5}); !chk(v, ok, 5, true) { pass = false }
	if v, ok := SafeAverage([]int64{-7,0}); !chk(v, ok, -3, true) { pass = false }
	if v, ok := SafeAverage([]int64{2,2,2,2}); !chk(v, ok, 2, true) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
