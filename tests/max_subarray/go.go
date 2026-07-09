import "fmt"
import "os"

func main() {
	ok := MaxSubarray([]int64{-2,1,-3,4,-1,2,1,-5,4}) == 6 &&
		MaxSubarray([]int64{}) == 0 &&
		MaxSubarray([]int64{-3,-1,-2}) == -1 &&
		MaxSubarray([]int64{5}) == 5
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
