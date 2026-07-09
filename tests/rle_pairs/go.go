import "fmt"
import "os"

func runsEq(a, b []Run) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	ok := runsEq(Rle([]int64{5,5,5,2}), []Run{{5,3},{2,1}}) &&
		runsEq(Rle([]int64{}), []Run{}) &&
		runsEq(Rle([]int64{1,2,2}), []Run{{1,1},{2,2}})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
