import "fmt"
import "os"

func main() {
	ok := AlternatingSum([]int64{10,3,2}) == 9 &&
		AlternatingSum([]int64{5,5}) == 0 &&
		AlternatingSum([]int64{7}) == 7 &&
		AlternatingSum([]int64{}) == 0 &&
		AlternatingSum([]int64{1,2,3,4}) == -2 &&
		AlternatingSum([]int64{-1,-2}) == 1
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
