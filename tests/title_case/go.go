import "fmt"
import "os"

func main() {
	ok := TitleCase("hello WORLD") == "Hello World" &&
		TitleCase("  a  b ") == "  A  B " &&
		TitleCase("3rd place") == "3rd Place" &&
		TitleCase("") == "" &&
		TitleCase("mIxEd CaSe") == "Mixed Case"
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
