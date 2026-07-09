pub fn interleave(a: &[i64], b: &[i64]) -> Vec<i64> {
    let mut out = Vec::new();
    let mut i = 0;
    while i < a.len() || i < b.len() {
        if i < a.len() { out.push(a[i]); }
        if i < b.len() { out.push(b[i]); }
        i += 1;
    }
    out
}
