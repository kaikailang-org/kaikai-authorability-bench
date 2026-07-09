import "fmt"
import "os"

func main() {
	ok := TankLevel(10, []int64{7,7,-20,3}) == 3 &&
		TankLevel(10, []int64{}) == 0 &&
		TankLevel(0, []int64{5,-2}) == 0 &&
		TankLevel(100, []int64{30,30}) == 60 &&
		TankLevel(5, []int64{9}) == 5 &&
		TankLevel(5, []int64{-3}) == 0
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
