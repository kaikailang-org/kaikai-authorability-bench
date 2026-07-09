import "fmt"
import "os"

func main() {
	ok := DayOfYear(2024, 3, 1) == 61 &&
		DayOfYear(2023, 3, 1) == 60 &&
		DayOfYear(2000, 1, 1) == 1 &&
		DayOfYear(1900, 3, 1) == 60 &&
		DayOfYear(2024, 12, 31) == 366 &&
		DayOfYear(2023, 12, 31) == 365
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
