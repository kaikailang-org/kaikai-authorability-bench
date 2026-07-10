import "fmt"
import "os"

func pEq(a []Param, b ...Param) bool {
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

	p1, o1 := ParseQuery("a=1&b=2")
	ok = ok && o1 && pEq(p1, Param{"a", "1"}, Param{"b", "2"})

	p2, o2 := ParseQuery("k=a=b")
	ok = ok && o2 && pEq(p2, Param{"k", "a=b"})

	p3, o3 := ParseQuery("a=")
	ok = ok && o3 && pEq(p3, Param{"a", ""})

	p4, o4 := ParseQuery("")
	ok = ok && o4 && pEq(p4)

	_, b1 := ParseQuery("=x")
	_, b2 := ParseQuery("a=1&&b=2")
	_, b3 := ParseQuery("flag")
	_, b4 := ParseQuery("a=1&flag")
	ok = ok && !b1 && !b2 && !b3 && !b4

	v, g1 := GetParam(p1, "b")
	ok = ok && g1 && v == "2"
	_, g2 := GetParam(p1, "z")
	ok = ok && !g2
	dup, _ := ParseQuery("x=1&x=2")
	v2, g3 := GetParam(dup, "x")
	ok = ok && g3 && v2 == "1"

	s1 := SetParam(p1, "a", "9")
	ok = ok && pEq(s1, Param{"a", "9"}, Param{"b", "2"})
	s2 := SetParam(p1, "c", "3")
	ok = ok && pEq(s2, Param{"a", "1"}, Param{"b", "2"}, Param{"c", "3"})
	s3 := SetParam(dup, "x", "7")
	ok = ok && pEq(s3, Param{"x", "7"}, Param{"x", "2"})
	// purity
	ok = ok && pEq(p1, Param{"a", "1"}, Param{"b", "2"})

	ok = ok && RenderQuery(p1) == "a=1&b=2"
	ok = ok && RenderQuery([]Param{}) == ""
	ok = ok && RenderQuery(p3) == "a="
	rt, _ := ParseQuery("k=a=b&z=")
	ok = ok && RenderQuery(rt) == "k=a=b&z="

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
