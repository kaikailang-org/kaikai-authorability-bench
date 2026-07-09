import "fmt"
import "os"

func rangesEq(a, b []Range) bool {
	if len(a) != len(b) { return false }
	for i := range a {
		if a[i].Start != b[i].Start || a[i].Finish != b[i].Finish { return false }
	}
	return true
}

func main() {
	ok := rangesEq(MergeRanges([]Range{{5,7},{1,3},{2,4}}), []Range{{1,4},{5,7}}) &&
		rangesEq(MergeRanges([]Range{{1,2},{2,3}}), []Range{{1,3}}) &&
		rangesEq(MergeRanges([]Range{}), []Range{}) &&
		rangesEq(MergeRanges([]Range{{1,10},{2,3}}), []Range{{1,10}}) &&
		rangesEq(MergeRanges([]Range{{4,5},{1,2}}), []Range{{1,2},{4,5}})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
