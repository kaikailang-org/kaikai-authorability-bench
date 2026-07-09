import "fmt"
import "os"

func chk(a, b Range, want Range) bool {
	r, ok := RangeOverlap(a, b)
	return ok && r == want
}

func noneOf(a, b Range) bool {
	_, ok := RangeOverlap(a, b)
	return !ok
}

func main() {
	ok := chk(Range{1,5}, Range{3,8}, Range{3,5}) &&
		noneOf(Range{1,3}, Range{3,5}) &&
		noneOf(Range{1,2}, Range{8,9}) &&
		chk(Range{2,9}, Range{3,4}, Range{3,4}) &&
		chk(Range{3,8}, Range{1,5}, Range{3,5})
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
