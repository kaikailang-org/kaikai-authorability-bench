import "fmt"
import "os"

func main() {
	ok := BalancedBrackets("([]{})") &&
		BalancedBrackets("f(a[0]) { return; }") &&
		!BalancedBrackets("([)]") &&
		!BalancedBrackets("(") &&
		BalancedBrackets("") &&
		!BalancedBrackets(")(") &&
		BalancedBrackets("no brackets at all")
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
