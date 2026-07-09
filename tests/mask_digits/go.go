import "fmt"
import "os"

func main() {
	ok := MaskDigits("1234 5678 9012 3456") == "**** **** **** 3456" &&
		MaskDigits("12-34-5") == "*2-34-5" &&
		MaskDigits("12") == "12" &&
		MaskDigits("") == "" &&
		MaskDigits("abcd") == "abcd" &&
		MaskDigits("12345") == "*2345"
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
