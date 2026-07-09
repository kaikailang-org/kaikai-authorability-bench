import "fmt"
import "os"

func main() {
	ok := RenderDuration(3725) == "1h02m05s" &&
		RenderDuration(90) == "1m30s" &&
		RenderDuration(5) == "5s" &&
		RenderDuration(0) == "0s" &&
		RenderDuration(3600) == "1h00m00s" &&
		RenderDuration(60) == "1m00s" &&
		RenderDuration(45296) == "12h34m56s"
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
