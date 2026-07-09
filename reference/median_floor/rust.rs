pub fn median_floor(xs: &[i64]) -> Option<i64> {
    if xs.is_empty() { return None; }
    let mut s = xs.to_vec();
    s.sort();
    let n = s.len();
    if n % 2 == 1 { Some(s[n / 2]) } else { Some((s[n / 2 - 1] + s[n / 2]) / 2) }
}
