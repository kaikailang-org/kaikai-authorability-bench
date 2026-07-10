import "fmt"
import "os"

func main() {
	ok := true

	l, o := ParseLine("WARN: disk low")
	ok = ok && o && l.Level == "WARN" && l.Message == "disk low"

	l2, o2 := ParseLine("INFO: ")
	ok = ok && o2 && l2.Level == "INFO" && l2.Message == ""

	l3, o3 := ParseLine("ERROR: a: b: c")
	ok = ok && o3 && l3.Level == "ERROR" && l3.Message == "a: b: c"

	_, b1 := ParseLine("INFO:x")
	_, b2 := ParseLine("DEBUG: x")
	_, b3 := ParseLine("info: x")
	_, b4 := ParseLine("INFO")
	_, b5 := ParseLine("")
	_, b6 := ParseLine(" INFO: x")
	ok = ok && !b1 && !b2 && !b3 && !b4 && !b5 && !b6

	lines := []string{"INFO: a", "ERROR: b", "INFO: c", "WARN: "}
	n, c1 := CountLevel(lines, "INFO")
	ok = ok && c1 && n == 2
	z, c2 := CountLevel([]string{}, "ERROR")
	ok = ok && c2 && z == 0
	_, c3 := CountLevel(lines, "DEBUG")
	_, c4 := CountLevel([]string{"INFO: a", "broken"}, "INFO")
	ok = ok && !c3 && !c4

	s, s1 := Summary(lines)
	ok = ok && s1 && s == "INFO=2 WARN=1 ERROR=1"
	s0, s2 := Summary([]string{})
	ok = ok && s2 && s0 == "INFO=0 WARN=0 ERROR=0"
	_, s3 := Summary([]string{"INFO: a", "oops: b"})
	ok = ok && !s3

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
