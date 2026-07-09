pub fn safe_average(xs: &[i64]) -> Option<i64> {
    if xs.is_empty() { return None; }
    Some(xs.iter().sum::<i64>() / xs.len() as i64)
}
