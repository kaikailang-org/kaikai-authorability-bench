import "fmt"
import "os"

func main() {
	ok := Caesar("abc", 1) == "bcd" &&
		Caesar("xyz", 3) == "abc" &&
		Caesar("abc", -1) == "zab" &&
		Caesar("Hello, World!", 13) == "Uryyb, Jbeyq!" &&
		Caesar("abc", 26) == "abc" &&
		Caesar("", 5) == ""
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
