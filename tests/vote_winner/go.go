import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	pass := true
	if v, ok := VoteWinner([]int64{3,1,3,2,3}); !chk(v, ok, 3, true) { pass = false }
	if v, ok := VoteWinner([]int64{1,2,1,2}); !chk(v, ok, 0, false) { pass = false }
	if v, ok := VoteWinner([]int64{}); !chk(v, ok, 0, false) { pass = false }
	if v, ok := VoteWinner([]int64{7}); !chk(v, ok, 7, true) { pass = false }
	if v, ok := VoteWinner([]int64{5,5,4}); !chk(v, ok, 5, true) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
