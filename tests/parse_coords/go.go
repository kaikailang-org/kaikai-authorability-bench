import "fmt"
import "os"

func chk(s string, wx, wy int64) bool {
	p, ok := ParseCoords(s)
	return ok && p.X == wx && p.Y == wy
}

func bad(s string) bool {
	_, ok := ParseCoords(s)
	return !ok
}

func main() {
	ok := chk("3,-4", 3, -4) &&
		chk("0,0", 0, 0) &&
		chk("-12,7", -12, 7) &&
		bad("3,") &&
		bad("3, 4") &&
		bad("3.5,1") &&
		bad("--2,1") &&
		bad("1,2,3") &&
		bad("")
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
