type Turtle struct{ X, Y, Dir int64 }

func TurtleNew() Turtle {
	return Turtle{0, 0, 0}
}

func Turn(t Turtle, r int64) Turtle {
	d := ((t.Dir+r)%4 + 4) % 4
	return Turtle{t.X, t.Y, d}
}

func Forward(t Turtle, n int64) Turtle {
	if n < 0 {
		n = 0
	}
	switch t.Dir {
	case 0:
		return Turtle{t.X, t.Y + n, t.Dir}
	case 1:
		return Turtle{t.X + n, t.Y, t.Dir}
	case 2:
		return Turtle{t.X, t.Y - n, t.Dir}
	default:
		return Turtle{t.X - n, t.Y, t.Dir}
	}
}

func Walk(cmds []string) (Turtle, bool) {
	t := TurtleNew()
	for _, c := range cmds {
		switch c {
		case "L":
			t = Turn(t, -1)
		case "R":
			t = Turn(t, 1)
		case "F":
			t = Forward(t, 1)
		default:
			return Turtle{}, false
		}
	}
	return t, true
}
