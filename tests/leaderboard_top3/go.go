import "fmt"
import "os"

func eqI64(a, b []int64) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func main() {
	ok := eqI64(LeaderboardTop3([]Score{{1, 50}, {2, 80}, {3, 50}, {4, 10}}), []int64{2, 1, 3}) &&
		eqI64(LeaderboardTop3([]Score{}), []int64{}) &&
		eqI64(LeaderboardTop3([]Score{{9, 1}}), []int64{9}) &&
		eqI64(LeaderboardTop3([]Score{{1, 5}, {2, 5}, {3, 5}, {4, 5}}), []int64{1, 2, 3}) &&
		eqI64(LeaderboardTop3([]Score{{1, -5}, {2, -1}}), []int64{2, 1})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
