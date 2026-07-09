import "strings"

func partVal(s string) (int64, bool) {
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
	return v, true
}

func SemverCompare(a, b string) (int64, bool) {
	pa := strings.Split(a, ".")
	pb := strings.Split(b, ".")
	if len(pa) != 3 || len(pb) != 3 {
		return 0, false
	}
	for i := 0; i < 3; i++ {
		va, oka := partVal(pa[i])
		vb, okb := partVal(pb[i])
		if !oka || !okb {
			return 0, false
		}
		if va < vb {
			return -1, true
		}
		if va > vb {
			return 1, true
		}
	}
	return 0, true
}
