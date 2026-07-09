import "fmt"
import "os"

func main() {
	ok := InvoiceTotal([]Line{{2, 250}, {1, 100}}) == 600 &&
		InvoiceTotal([]Line{}) == 0 &&
		InvoiceTotal([]Line{{0, 999}}) == 0 &&
		InvoiceTotal([]Line{{-3, 100}, {4, 25}}) == 100
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
