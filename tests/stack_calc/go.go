import "fmt"
import "os"

func sEq(a []int64, b ...int64) bool {
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

func runOK(prog []string, want int64) bool {
	v, o := Run(prog)
	return o && v == want
}

func runBad(prog []string) bool {
	_, o := Run(prog)
	return !o
}

func main() {
	ok := sEq(Push([]int64{2, 1}, 3), 3, 2, 1)

	s1, o1 := Step([]int64{3, 2}, "add")
	ok = ok && o1 && sEq(s1, 5)
	s2, o2 := Step([]int64{3, 2, 9}, "mul")
	ok = ok && o2 && sEq(s2, 6, 9)
	s3, o3 := Step([]int64{1, 2}, "swap")
	ok = ok && o3 && sEq(s3, 2, 1)
	s4, o4 := Step([]int64{7}, "dup")
	ok = ok && o4 && sEq(s4, 7, 7)
	s5, o5 := Step([]int64{7, 8}, "pop")
	ok = ok && o5 && sEq(s5, 8)
	_, b1 := Step([]int64{1}, "add")
	_, b2 := Step([]int64{}, "dup")
	_, b3 := Step([]int64{1, 2}, "sub")
	ok = ok && !b1 && !b2 && !b3

	ok = ok && runOK([]string{"2", "3", "add"}, 5) &&
		runOK([]string{"2", "3", "mul", "dup", "add"}, 12) &&
		runOK([]string{"07"}, 7) &&
		runOK([]string{"-4", "2", "add"}, -2) &&
		runOK([]string{"1", "2", "swap", "pop"}, 2) &&
		runBad([]string{"4", "pop"}) &&
		runBad([]string{"swap"}) &&
		runBad([]string{"-"}) &&
		runBad([]string{}) &&
		runBad([]string{"1", "2"}) &&
		runBad([]string{"+5"})

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
