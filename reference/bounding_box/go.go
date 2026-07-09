type Point struct{ X, Y int64 }
type Box struct{ MinX, MinY, MaxX, MaxY int64 }

func BoundingBox(ps []Point) (Box, bool) {
	if len(ps) == 0 {
		return Box{}, false
	}
	b := Box{ps[0].X, ps[0].Y, ps[0].X, ps[0].Y}
	for _, p := range ps {
		if p.X < b.MinX { b.MinX = p.X }
		if p.Y < b.MinY { b.MinY = p.Y }
		if p.X > b.MaxX { b.MaxX = p.X }
		if p.Y > b.MaxY { b.MaxY = p.Y }
	}
	return b, true
}
