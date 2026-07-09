pub fn debounce(xs: &[i64]) -> Vec<i64> {
    let mut out: Vec<i64> = Vec::new();
    for &x in xs {
        if out.last() != Some(&x) { out.push(x); }
    }
    out
}
