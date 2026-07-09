import "fmt"
import "os"

func main() {
	ok := Slugify("Hello, World!") == "hello-world" &&
		Slugify("  42 is  GREAT ") == "42-is-great" &&
		Slugify("---") == "" &&
		Slugify("") == "" &&
		Slugify("already-good") == "already-good" &&
		Slugify("A") == "a"
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
