import "strings"

func ColWidths(headers []string, rows [][]string) ([]int64, bool) {
	if len(headers) == 0 {
		return nil, false
	}
	ws := make([]int64, len(headers))
	for i, h := range headers {
		ws[i] = int64(len(h))
	}
	for _, r := range rows {
		if len(r) != len(headers) {
			return nil, false
		}
		for i, c := range r {
			if int64(len(c)) > ws[i] {
				ws[i] = int64(len(c))
			}
		}
	}
	return ws, true
}

func RenderRow(cells []string, widths []int64) string {
	parts := make([]string, len(cells))
	for i, c := range cells {
		if i < len(cells)-1 {
			parts[i] = c + strings.Repeat(" ", int(widths[i])-len(c))
		} else {
			parts[i] = c
		}
	}
	return strings.Join(parts, " | ")
}

func RenderTable(headers []string, rows [][]string) (string, bool) {
	ws, ok := ColWidths(headers, rows)
	if !ok {
		return "", false
	}
	dashes := make([]string, len(ws))
	for i, w := range ws {
		dashes[i] = strings.Repeat("-", int(w))
	}
	lines := []string{RenderRow(headers, ws), strings.Join(dashes, " | ")}
	for _, r := range rows {
		lines = append(lines, RenderRow(r, ws))
	}
	return strings.Join(lines, "\n"), true
}
