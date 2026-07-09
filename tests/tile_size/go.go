import "fmt"
import "os"

func main() {
	ok := TileSize(12, 18) == 6 &&
		TileSize(7, 7) == 7 &&
		TileSize(13, 7) == 1 &&
		TileSize(100, 10) == 10 &&
		TileSize(1, 999) == 1 &&
		TileSize(18, 12) == 6
	if !ok { fmt.Println("FAIL"); os.Exit(1) }
	fmt.Println("OK")
}
