import "fmt"
import "os"

func main() {
	ok := NextFreeID([]int64{0,1,3}) == 2 &&
		NextFreeID([]int64{}) == 0 &&
		NextFreeID([]int64{1,2,3}) == 0 &&
		NextFreeID([]int64{0,0,1,1,2}) == 3 &&
		NextFreeID([]int64{3,1,0,2}) == 4
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
