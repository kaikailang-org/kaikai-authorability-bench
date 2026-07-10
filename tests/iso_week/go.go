import "fmt"
import "os"

func chk(y, m, d, wy, ww int64) bool {
	r := IsoWeekOf(y, m, d)
	return r.Year == wy && r.Week == ww
}

func main() {
	ok := chk(2026, 1, 1, 2026, 1) &&
		chk(2016, 1, 1, 2015, 53) &&
		chk(2025, 12, 29, 2026, 1) &&
		chk(2021, 1, 3, 2020, 53) &&
		chk(2020, 2, 29, 2020, 9) &&
		chk(2027, 1, 1, 2026, 53) &&
		chk(2020, 12, 31, 2020, 53) &&
		chk(1600, 1, 1, 1599, 52) &&
		chk(2000, 1, 1, 1999, 52) &&
		chk(2024, 12, 30, 2025, 1) &&
		chk(1999, 12, 31, 1999, 52) &&
		chk(2015, 12, 28, 2015, 53) &&
		chk(2026, 7, 9, 2026, 28)
	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
