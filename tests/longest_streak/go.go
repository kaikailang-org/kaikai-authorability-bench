import "fmt"
import "os"

func main() {
	ok := LongestStreak([]bool{true,true,false,true,true,true}) == 3 &&
		LongestStreak([]bool{}) == 0 &&
		LongestStreak([]bool{false,false}) == 0 &&
		LongestStreak([]bool{true,true,true,true}) == 4 &&
		LongestStreak([]bool{true,false,true}) == 1
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
