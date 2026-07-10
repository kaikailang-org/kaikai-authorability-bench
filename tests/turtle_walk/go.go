import "fmt"
import "os"

func at(t Turtle, x, y, d int64) bool {
	return t.X == x && t.Y == y && t.Dir == d
}

func main() {
	ok := at(TurtleNew(), 0, 0, 0)

	ok = ok && Turn(TurtleNew(), -1).Dir == 3
	ok = ok && Turn(TurtleNew(), 6).Dir == 2
	ok = ok && Turn(TurtleNew(), -7).Dir == 1
	ok = ok && Turn(Turtle{5, 5, 3}, 1).Dir == 0

	ok = ok && at(Forward(TurtleNew(), 3), 0, 3, 0)
	ok = ok && at(Forward(Turtle{0, 0, 2}, 2), 0, -2, 2)
	ok = ok && at(Forward(Turtle{1, 1, 3}, 4), -3, 1, 3)
	ok = ok && at(Forward(Turtle{9, 9, 1}, -5), 9, 9, 1)

	w, o := Walk([]string{"F", "R", "F", "F"})
	ok = ok && o && at(w, 2, 1, 1)
	w2, o2 := Walk([]string{"L", "F", "F", "L", "F"})
	ok = ok && o2 && at(w2, -2, -1, 2)
	w3, o3 := Walk([]string{})
	ok = ok && o3 && at(w3, 0, 0, 0)
	_, o4 := Walk([]string{"F", "X"})
	ok = ok && !o4
	_, o5 := Walk([]string{"f"})
	ok = ok && !o5

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
