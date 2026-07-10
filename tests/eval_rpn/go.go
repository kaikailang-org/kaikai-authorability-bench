import "fmt"
import "os"

func chk(tokens []string, want int64, ok bool) bool {
	v, o := EvalRPN(tokens)
	if !ok {
		return !o
	}
	return o && v == want
}

func main() {
	ok := chk([]string{"3", "4", "+"}, 7, true) &&
		chk([]string{"5", "1", "2", "+", "4", "*", "-"}, -7, true) &&
		chk([]string{"7", "-2", "/"}, -3, true) &&
		chk([]string{"-7", "2", "/"}, -3, true) &&
		chk([]string{"42"}, 42, true) &&
		chk([]string{"2", "3", "4", "*", "+"}, 14, true) &&
		chk([]string{"8", "0", "/"}, 0, false) &&
		chk([]string{"3", "4"}, 0, false) &&
		chk([]string{"+"}, 0, false) &&
		chk([]string{"3", "-"}, 0, false) &&
		chk([]string{}, 0, false) &&
		chk([]string{"3", "x", "+"}, 0, false) &&
		chk([]string{"-"}, 0, false) &&
		chk([]string{"+5", "1", "+"}, 0, false) &&
		chk([]string{"0", "5", "/"}, 0, true)
	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
