import "fmt"
import "os"

func main() {
	ok := CollatzSteps(1) == 0 && CollatzSteps(2) == 1 &&
		CollatzSteps(3) == 7 && CollatzSteps(27) == 111
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
