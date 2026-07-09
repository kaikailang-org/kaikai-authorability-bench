import "fmt"
import "os"

func main() {
	ok := CleanPalindrome("A man, a plan, a canal: Panama") &&
		!CleanPalindrome("race a car") &&
		CleanPalindrome("") &&
		CleanPalindrome("12 21") &&
		CleanPalindrome("No 'x' in Nixon") &&
		!CleanPalindrome("1a2")
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
