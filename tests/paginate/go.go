import "fmt"
import "os"

func main() {
	ok := Paginate(10, 3) == Paging{4, 1} &&
		Paginate(9, 3) == Paging{3, 3} &&
		Paginate(0, 5) == Paging{0, 0} &&
		Paginate(1, 1) == Paging{1, 1} &&
		Paginate(5, 10) == Paging{1, 5}
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
