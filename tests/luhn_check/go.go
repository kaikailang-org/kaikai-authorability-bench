import "fmt"
import "os"

func main() {
	ok := LuhnCheck([]int64{7,9,9,2,7,3,9,8,7,1,3}) &&
		!LuhnCheck([]int64{7,9,9,2,7,3,9,8,7,1,1}) &&
		!LuhnCheck([]int64{}) &&
		LuhnCheck([]int64{0})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
