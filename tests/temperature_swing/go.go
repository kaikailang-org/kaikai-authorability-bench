import "fmt"
import "os"

func main() {
	ok := TemperatureSwing([]int64{10,13,9}) == 4 &&
		TemperatureSwing([]int64{5}) == 0 &&
		TemperatureSwing([]int64{}) == 0 &&
		TemperatureSwing([]int64{-3,4}) == 7 &&
		TemperatureSwing([]int64{1,1,1}) == 0 &&
		TemperatureSwing([]int64{0,-10,5}) == 15
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
