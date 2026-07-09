import "fmt"
import "os"

func main() {
	ok := RunningMinBalance([]int64{100,-30,-90,50}) == -20 &&
		RunningMinBalance([]int64{5,10}) == 0 &&
		RunningMinBalance([]int64{}) == 0 &&
		RunningMinBalance([]int64{-1,-1,-1}) == -3 &&
		RunningMinBalance([]int64{-5,10,-3}) == -5
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
