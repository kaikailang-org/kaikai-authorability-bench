import "fmt"
import "os"

func main() {
	ok := IntSqrt(0) == 0 &&
		IntSqrt(1) == 1 &&
		IntSqrt(8) == 2 &&
		IntSqrt(9) == 3 &&
		IntSqrt(999999999999) == 999999 &&
		IntSqrt(1000000000000) == 1000000 &&
		IntSqrt(2) == 1
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
