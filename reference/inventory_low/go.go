type Stock struct{ Sku, Qty int64 }

func InventoryLow(stock []Stock, threshold int64) []int64 {
	out := []int64{}
	for _, s := range stock {
		if s.Qty < threshold {
			out = append(out, s.Sku)
		}
	}
	return out
}
