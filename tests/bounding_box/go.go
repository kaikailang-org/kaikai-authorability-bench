import "fmt"
import "os"

func main() {
	b, ok := BoundingBox([]Point{{1, 5}, {-2, 3}, {4, 4}})
	pass := ok && b == Box{-2, 3, 4, 5}
	b2, ok2 := BoundingBox([]Point{{7, -1}})
	pass = pass && ok2 && b2 == Box{7, -1, 7, -1}
	_, ok3 := BoundingBox([]Point{})
	pass = pass && !ok3
	if !pass { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
