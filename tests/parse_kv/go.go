import "fmt"
import "os"

func good(s string, want []Entry) bool {
	es, ok := ParseKV(s)
	if !ok || len(es) != len(want) { return false }
	for i := range es {
		if es[i] != want[i] { return false }
	}
	return true
}

func bad(s string) bool {
	_, ok := ParseKV(s)
	return !ok
}

func main() {
	pass := good("a=1;b=2", []Entry{{"a", "1"}, {"b", "2"}}) &&
		good("a=", []Entry{{"a", ""}}) &&
		good("", []Entry{}) &&
		good("host=localhost", []Entry{{"host", "localhost"}}) &&
		bad("=1") &&
		bad("a") &&
		bad("a=1;") &&
		bad("a=b=c")
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
