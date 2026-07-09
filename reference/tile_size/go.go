func TileSize(w, h int64) int64 {
	for h != 0 {
		w, h = h, w%h
	}
	return w
}
