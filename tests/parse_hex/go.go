import "fmt"
import "os"

func chk(got int64, ok bool, want int64, wantOk bool) bool {
	if ok != wantOk { return false }
	return !ok || got == want
}

func main() {
	pass := true
	if v, ok := ParseHex("0x1F"); !chk(v, ok, 31, true) { pass = false }
	if v, ok := ParseHex("0xff"); !chk(v, ok, 255, true) { pass = false }
	if v, ok := ParseHex("0x0"); !chk(v, ok, 0, true) { pass = false }
	if v, ok := ParseHex("0xDeadBeef"); !chk(v, ok, 3735928559, true) { pass = false }
	if v, ok := ParseHex("1F"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := ParseHex("0x"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := ParseHex("0xG1"); !chk(v, ok, 0, false) { pass = false }
	if v, ok := ParseHex(""); !chk(v, ok, 0, false) { pass = false }
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
