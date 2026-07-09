import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	items := []Item{{10, 250}, {7, 100}, {10, 999}}
	pass := true
	if v, ok := LookupPrice(items, 7); !chk(v, ok, 100, true) { pass = false }
	if v, ok := LookupPrice(items, 10); !chk(v, ok, 250, true) { pass = false }
	if v, ok := LookupPrice(items, 3); !chk(v, ok, 0, false) { pass = false }
	if v, ok := LookupPrice([]Item{}, 1); !chk(v, ok, 0, false) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
