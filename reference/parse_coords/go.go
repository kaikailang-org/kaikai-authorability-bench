import "strings"

type Point struct{ X, Y int64 }

func intVal(s string) (int64, bool) {
	neg := false
	if strings.HasPrefix(s, "-") {
		neg = true
		s = s[1:]
	}
	if len(s) == 0 {
		return 0, false
	}
	var v int64
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		v = v*10 + int64(c-'0')
	}
	if neg {
		v = -v
	}
	return v, true
}

func ParseCoords(s string) (Point, bool) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return Point{}, false
	}
	x, okx := intVal(parts[0])
	y, oky := intVal(parts[1])
	if !okx || !oky {
		return Point{}, false
	}
	return Point{X: x, Y: y}, true
}
