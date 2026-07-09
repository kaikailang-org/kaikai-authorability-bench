import "fmt"
import "os"

func main() {
	ok := Initials("grace brewster hopper") == "GBH" &&
		Initials(" ada   lovelace ") == "AL" &&
		Initials("") == "" &&
		Initials("x") == "X" &&
		Initials("Alan Turing") == "AT"
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
