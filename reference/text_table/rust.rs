pub fn col_widths(headers: &[&str], rows: &[Vec<&str>]) -> Option<Vec<i64>> {
    if headers.is_empty() {
        return None;
    }
    let mut ws: Vec<i64> = headers.iter().map(|h| h.len() as i64).collect();
    for r in rows {
        if r.len() != headers.len() {
            return None;
        }
        for (i, c) in r.iter().enumerate() {
            ws[i] = ws[i].max(c.len() as i64);
        }
    }
    Some(ws)
}

pub fn render_row(cells: &[&str], widths: &[i64]) -> String {
    cells
        .iter()
        .enumerate()
        .map(|(i, c)| {
            if i < cells.len() - 1 {
                format!("{}{}", c, " ".repeat(widths[i] as usize - c.len()))
            } else {
                c.to_string()
            }
        })
        .collect::<Vec<_>>()
        .join(" | ")
}

pub fn render_table(headers: &[&str], rows: &[Vec<&str>]) -> Option<String> {
    let ws = col_widths(headers, rows)?;
    let dashes: Vec<String> = ws.iter().map(|w| "-".repeat(*w as usize)).collect();
    let mut lines = vec![
        render_row(headers, &ws),
        dashes.join(" | "),
    ];
    for r in rows {
        lines.push(render_row(r, &ws));
    }
    Some(lines.join("\n"))
}
