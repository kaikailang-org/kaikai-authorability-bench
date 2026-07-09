fn part_val(s: &str) -> Option<i64> {
    if s.is_empty() || !s.chars().all(|c| c.is_ascii_digit()) { return None; }
    let mut v: i64 = 0;
    for c in s.chars() { v = v * 10 + (c as i64 - '0' as i64); }
    Some(v)
}

pub fn semver_compare(a: &str, b: &str) -> Option<i64> {
    let pa = a.split('.').map(part_val).collect::<Option<Vec<_>>>()?;
    let pb = b.split('.').map(part_val).collect::<Option<Vec<_>>>()?;
    if pa.len() != 3 || pb.len() != 3 { return None; }
    for i in 0..3 {
        if pa[i] < pb[i] { return Some(-1); }
        if pa[i] > pb[i] { return Some(1); }
    }
    Some(0)
}
