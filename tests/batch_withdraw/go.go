import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	pass := true
	if v, ok := BatchWithdraw(100, []int64{30,50}); !chk(v, ok, 20, true) { pass = false }
	if v, ok := BatchWithdraw(100, []int64{30,80}); !chk(v, ok, 0, false) { pass = false }
	if v, ok := BatchWithdraw(100, []int64{}); !chk(v, ok, 100, true) { pass = false }
	if v, ok := BatchWithdraw(100, []int64{100}); !chk(v, ok, 0, true) { pass = false }
	if v, ok := BatchWithdraw(100, []int64{0}); !chk(v, ok, 0, false) { pass = false }
	if v, ok := BatchWithdraw(100, []int64{-5}); !chk(v, ok, 0, false) { pass = false }
	if v, ok := BatchWithdraw(0, []int64{1}); !chk(v, ok, 0, false) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
