import "fmt"
import "os"

func chk(s, want string, ok bool) bool {
	v, o := NormalizePath(s)
	if !ok {
		return !o
	}
	return o && v == want
}

func main() {
	ok := chk("/a//b/./c/", "/a/b/c", true) &&
		chk("/a/../../b", "/b", true) &&
		chk("/../", "/", true) &&
		chk("/...", "/...", true) &&
		chk("/", "/", true) &&
		chk("", "", false) &&
		chk("a/b", "", false) &&
		chk("/a/b/../c", "/a/c", true) &&
		chk("/a/./b/..", "/a", true) &&
		chk("/..hidden", "/..hidden", true) &&
		chk("/a.b/c.", "/a.b/c.", true) &&
		chk("//", "/", true) &&
		chk("/a/../a/../a", "/a", true) &&
		chk(".", "", false)
	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
