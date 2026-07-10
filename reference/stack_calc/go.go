import "strconv"

func Push(st []int64, v int64) []int64 {
	return append([]int64{v}, st...)
}

func litOK(t string) bool {
	i := 0
	if len(t) > 0 && t[0] == '-' {
		i = 1
	}
	if i == len(t) {
		return false
	}
	for ; i < len(t); i++ {
		if t[i] < '0' || t[i] > '9' {
			return false
		}
	}
	return true
}

func Step(st []int64, op string) ([]int64, bool) {
	switch op {
	case "add", "mul":
		if len(st) < 2 {
			return nil, false
		}
		b, a := st[0], st[1]
		r := a + b
		if op == "mul" {
			r = a * b
		}
		return append([]int64{r}, st[2:]...), true
	case "swap":
		if len(st) < 2 {
			return nil, false
		}
		return append([]int64{st[1], st[0]}, st[2:]...), true
	case "dup":
		if len(st) < 1 {
			return nil, false
		}
		return append([]int64{st[0]}, st...), true
	case "pop":
		if len(st) < 1 {
			return nil, false
		}
		return append([]int64{}, st[1:]...), true
	default:
		if !litOK(op) {
			return nil, false
		}
		v, err := strconv.ParseInt(op, 10, 64)
		if err != nil {
			return nil, false
		}
		return Push(st, v), true
	}
}

func Run(prog []string) (int64, bool) {
	st := []int64{}
	for _, op := range prog {
		next, ok := Step(st, op)
		if !ok {
			return 0, false
		}
		st = next
	}
	if len(st) != 1 {
		return 0, false
	}
	return st[0], true
}
