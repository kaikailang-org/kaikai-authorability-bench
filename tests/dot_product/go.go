import "fmt"
import "os"

func main() {
	r1, ok1 := DotProduct([]int64{1,2,3}, []int64{4,5,6})
	_, ok2 := DotProduct([]int64{1}, []int64{})
	r3, ok3 := DotProduct([]int64{}, []int64{})
	ok := ok1 && r1 == 32 && !ok2 && ok3 && r3 == 0
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
