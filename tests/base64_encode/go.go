import "fmt"
import "os"

func chk(bs []int64, want string, ok bool) bool {
	v, o := Base64Encode(bs)
	if !ok {
		return !o
	}
	return o && v == want
}

func main() {
	ok := chk([]int64{77, 97, 110}, "TWFu", true) &&
		chk([]int64{77, 97}, "TWE=", true) &&
		chk([]int64{77}, "TQ==", true) &&
		chk([]int64{102, 111, 111, 98}, "Zm9vYg==", true) &&
		chk([]int64{102, 111, 111, 98, 97}, "Zm9vYmE=", true) &&
		chk([]int64{102, 111, 111, 98, 97, 114}, "Zm9vYmFy", true) &&
		chk([]int64{}, "", true) &&
		chk([]int64{0}, "AA==", true) &&
		chk([]int64{0, 0, 0}, "AAAA", true) &&
		chk([]int64{255, 255, 255}, "////", true) &&
		chk([]int64{251}, "+w==", true) &&
		chk([]int64{256}, "", false) &&
		chk([]int64{-1}, "", false) &&
		chk([]int64{77, 300, 97}, "", false)
	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
