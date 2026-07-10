type Item struct{ Sku, Qty, Reserved int64 }

func find(inv []Item, sku int64) int {
	for i, it := range inv {
		if it.Sku == sku {
			return i
		}
	}
	return -1
}

func clone(inv []Item) []Item {
	return append([]Item{}, inv...)
}

func AddStock(inv []Item, sku, n int64) ([]Item, bool) {
	if n < 1 {
		return nil, false
	}
	out := clone(inv)
	if i := find(out, sku); i >= 0 {
		out[i].Qty += n
	} else {
		out = append(out, Item{sku, n, 0})
	}
	return out, true
}

func Reserve(inv []Item, sku, n int64) ([]Item, bool) {
	i := find(inv, sku)
	if n < 1 || i < 0 || inv[i].Qty-inv[i].Reserved < n {
		return nil, false
	}
	out := clone(inv)
	out[i].Reserved += n
	return out, true
}

func Release(inv []Item, sku, n int64) ([]Item, bool) {
	i := find(inv, sku)
	if n < 1 || i < 0 || inv[i].Reserved < n {
		return nil, false
	}
	out := clone(inv)
	out[i].Reserved -= n
	return out, true
}

func Ship(inv []Item, sku, n int64) ([]Item, bool) {
	i := find(inv, sku)
	if n < 1 || i < 0 || inv[i].Reserved < n {
		return nil, false
	}
	out := clone(inv)
	out[i].Qty -= n
	out[i].Reserved -= n
	return out, true
}

func Available(inv []Item, sku int64) (int64, bool) {
	i := find(inv, sku)
	if i < 0 {
		return 0, false
	}
	return inv[i].Qty - inv[i].Reserved, true
}
