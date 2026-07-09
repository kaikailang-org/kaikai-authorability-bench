import "fmt"
import "os"

func main() {
	ok := DigitalRoot(0) == 0 &&
		DigitalRoot(9) == 9 &&
		DigitalRoot(942) == 6 &&
		DigitalRoot(132189) == 6 &&
		DigitalRoot(10) == 1 &&
		DigitalRoot(999999999) == 9
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
