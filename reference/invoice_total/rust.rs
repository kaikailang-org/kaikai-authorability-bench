pub struct Line { pub qty: i64, pub unit_cents: i64 }

pub fn invoice_total(lines: &[Line]) -> i64 {
    lines.iter().filter(|l| l.qty > 0).map(|l| l.qty * l.unit_cents).sum()
}
