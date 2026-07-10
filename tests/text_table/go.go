import "fmt"
import "os"

func wEq(a []int64, b ...int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	ok := true

	ws, o := ColWidths([]string{"name", "qty"}, [][]string{{"ana", "3"}, {"b", "12"}})
	ok = ok && o && wEq(ws, 4, 3)

	want := "name | qty\n---- | ---\nana  | 3\nb    | 12"
	got, o2 := RenderTable([]string{"name", "qty"}, [][]string{{"ana", "3"}, {"b", "12"}})
	ok = ok && o2 && got == want

	// header-only table
	got2, o3 := RenderTable([]string{"id"}, [][]string{})
	ok = ok && o3 && got2 == "id\n--"

	// cell wider than header drives the width
	got3, o4 := RenderTable([]string{"h"}, [][]string{{"wide"}})
	ok = ok && o4 && got3 == "h\n----\nwide"

	// last cell unpadded even when short
	ok = ok && RenderRow([]string{"a", "b"}, []int64{3, 5}) == "a   | b"
	// single column: no separator, no padding
	ok = ok && RenderRow([]string{"x"}, []int64{7}) == "x"

	// invalid shapes
	_, b1 := ColWidths([]string{}, [][]string{})
	_, b2 := ColWidths([]string{"a"}, [][]string{{"x", "y"}})
	_, b3 := RenderTable([]string{"a", "b"}, [][]string{{"only"}})
	ok = ok && !b1 && !b2 && !b3

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
