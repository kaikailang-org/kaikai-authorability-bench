type Line struct{ Qty, UnitCents int64 }

func InvoiceTotal(lines []Line) int64 {
	var total int64
	for _, l := range lines {
		if l.Qty > 0 {
			total += l.Qty * l.UnitCents
		}
	}
	return total
}
