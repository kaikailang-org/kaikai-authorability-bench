import "fmt"
import "os"

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func chk(s string, want []string, ok bool) bool {
	v, o := ParseCsvRow(s)
	if !ok {
		return !o
	}
	return o && eq(v, want)
}

func main() {
	ok := chk("a,b,c", []string{"a", "b", "c"}, true) &&
		chk("\"a,b\",c", []string{"a,b", "c"}, true) &&
		chk("a,\"x\"\"y\"", []string{"a", "x\"y"}, true) &&
		chk("\"\"", []string{""}, true) &&
		chk("", []string{""}, true) &&
		chk("a,", []string{"a", ""}, true) &&
		chk(",a", []string{"", "a"}, true) &&
		chk(",,", []string{"", "", ""}, true) &&
		chk("\"\"\"\"", []string{"\""}, true) &&
		chk("\"a\nb\",c", []string{"a\nb", "c"}, true) &&
		chk("a,\"b\"x", nil, false) &&
		chk("a\"b", nil, false) &&
		chk("\"a", nil, false) &&
		chk("\"a\"\"", nil, false) &&
		chk("x,\"a\"\"", nil, false)
	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
