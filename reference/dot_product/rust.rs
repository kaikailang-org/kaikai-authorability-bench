pub fn dot_product(a: &[i64], b: &[i64]) -> Option<i64> {
    if a.len() != b.len() { return None; }
    Some(a.iter().zip(b.iter()).map(|(x, y)| x * y).sum())
}
