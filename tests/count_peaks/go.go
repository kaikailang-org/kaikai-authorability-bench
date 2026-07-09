import "fmt"
import "os"

func main() {
	ok := CountPeaks([]int64{0,5,0,5,0}) == 2 &&
		CountPeaks([]int64{1,3,2}) == 1 &&
		CountPeaks([]int64{1,2,3}) == 0 &&
		CountPeaks([]int64{}) == 0 &&
		CountPeaks([]int64{1,2}) == 0 &&
		CountPeaks([]int64{1,3,3,1}) == 0
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
