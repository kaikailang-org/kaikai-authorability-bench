import "fmt"
import "os"

func main() {
	ok := CountWords("  hello   world ") == 2 &&
		CountWords("") == 0 &&
		CountWords("   ") == 0 &&
		CountWords("one") == 1 &&
		CountWords("a\tb") == 2 &&
		CountWords("a b c d") == 4
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
