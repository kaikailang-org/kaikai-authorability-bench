import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	pass := true
	if v, ok := SemverCompare("1.2.3", "1.2.3"); !chk(v, ok, 0, true) { pass = false }
	if v, ok := SemverCompare("1.2.3", "1.10.0"); !chk(v, ok, -1, true) { pass = false }
	if v, ok := SemverCompare("2.0.0", "1.9.9"); !chk(v, ok, 1, true) { pass = false }
	if v, ok := SemverCompare("1.2", "1.2.3"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := SemverCompare("1.2.3", "1.2.x"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := SemverCompare("1.2.3.4", "1.2.3"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := SemverCompare("1.-2.3", "1.2.3"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := SemverCompare("01.2.3", "1.2.3"); !chk(v, ok, 0, true) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
