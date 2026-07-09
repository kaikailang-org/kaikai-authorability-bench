type Item struct{ Sku, Cents int64 }

func LookupPrice(items []Item, sku int64) (int64, bool) {
	for _, i := range items {
		if i.Sku == sku {
			return i.Cents, true
		}
	}
	return 0, false
}
