import "fmt"
import "os"

func add(book []Student, id, score int64) []Student {
	out, _ := AddScore(book, id, score)
	return out
}

func avIs(book []Student, id, want int64) bool {
	v, o := AverageFloor(book, id)
	return o && v == want
}

func main() {
	ok := true

	b := add([]Student{}, 1, 90)
	b = add(b, 1, 95)
	ok = ok && len(b) == 1 && avIs(b, 1, 92)

	b = add(b, 1, 92)
	ok = ok && avIs(b, 1, 92) // 277/3 floors to 92

	b = add(b, 2, 80)
	ok = ok && len(b) == 2 && b[0].Id == 1 && b[1].Id == 2

	top, o := TopStudent(b)
	ok = ok && o && top == 1

	// tie at the floor -> no top student
	t := add(add([]Student{}, 1, 92), 2, 92)
	_, o2 := TopStudent(t)
	ok = ok && !o2

	// students without scores are ignored... but AddScore always scores.
	// an empty book has no top student
	_, o3 := TopStudent([]Student{})
	ok = ok && !o3

	// no-score student ignored by top, invalid for average
	nb := []Student{{7, []int64{}}, {8, []int64{50}}}
	topN, o4 := TopStudent(nb)
	ok = ok && o4 && topN == 8
	_, o5 := AverageFloor(nb, 7)
	ok = ok && !o5
	_, o6 := AverageFloor(nb, 99)
	ok = ok && !o6

	// score bounds
	_, v1 := AddScore(b, 1, 101)
	_, v2 := AddScore(b, 1, -1)
	ok = ok && !v1 && !v2
	b2, v3 := AddScore(b, 3, 0)
	ok = ok && v3 && avIs(b2, 3, 0)
	b3, v4 := AddScore(b, 3, 100)
	ok = ok && v4 && avIs(b3, 3, 100)

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
