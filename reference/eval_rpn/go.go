import "strconv"

func validLit(t string) bool {
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

func EvalRPN(tokens []string) (int64, bool) {
	st := []int64{}
	for _, t := range tokens {
		switch t {
		case "+", "-", "*", "/":
			n := len(st)
			if n < 2 {
				return 0, false
			}
			a, b := st[n-2], st[n-1]
			st = st[:n-2]
			var r int64
			switch t {
			case "+":
				r = a + b
			case "-":
				r = a - b
			case "*":
				r = a * b
			case "/":
				if b == 0 {
					return 0, false
				}
				r = a / b
			}
			st = append(st, r)
		default:
			if !validLit(t) {
				return 0, false
			}
			v, err := strconv.ParseInt(t, 10, 64)
			if err != nil {
				return 0, false
			}
			st = append(st, v)
		}
	}
	if len(st) != 1 {
		return 0, false
	}
	return st[0], true
}
