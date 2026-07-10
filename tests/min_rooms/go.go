import "fmt"
import "os"

func chk(ms []Meeting, want int64, ok bool) bool {
	v, o := MinRooms(ms)
	if !ok {
		return !o
	}
	return o && v == want
}

func main() {
	ok := chk([]Meeting{{0, 30}, {5, 10}, {15, 20}}, 2, true) &&
		chk([]Meeting{{1, 3}, {3, 5}}, 1, true) &&
		chk([]Meeting{{7, 10}, {2, 4}}, 1, true) &&
		chk([]Meeting{{1, 5}, {2, 5}, {3, 5}}, 3, true) &&
		chk([]Meeting{}, 0, true) &&
		chk([]Meeting{{4, 9}}, 1, true) &&
		chk([]Meeting{{1, 10}, {2, 3}, {4, 5}, {6, 7}}, 2, true) &&
		chk([]Meeting{{0, 5}, {5, 10}, {5, 8}}, 2, true) &&
		chk([]Meeting{{5, 5}}, 0, false) &&
		chk([]Meeting{{1, 3}, {6, 2}}, 0, false) &&
		chk([]Meeting{{-5, -1}, {-3, 2}}, 2, true)
	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
