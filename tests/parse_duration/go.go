import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	pass := true
	if v, ok := ParseDuration("2h30m15s"); !chk(v, ok, 9015, true) { pass = false }
	if v, ok := ParseDuration("90s"); !chk(v, ok, 90, true) { pass = false }
	if v, ok := ParseDuration("2h"); !chk(v, ok, 7200, true) { pass = false }
	if v, ok := ParseDuration("1h0m5s"); !chk(v, ok, 3605, true) { pass = false }
	if v, ok := ParseDuration("90m"); !chk(v, ok, 5400, true) { pass = false }
	if v, ok := ParseDuration(""); !chk(v, ok, 0, false) { pass = false }
	if v, ok := ParseDuration("5x"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := ParseDuration("m5"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := ParseDuration("1s1h"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := ParseDuration("2h30"); !chk(v, ok, 0, false) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
