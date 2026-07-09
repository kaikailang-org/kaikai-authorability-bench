pub fn histogram_buckets(xs: &[i64]) -> Vec<i64> {
    let mut out = vec![0i64; 10];
    for &x in xs {
        out[(x / 10) as usize] += 1;
    }
    out
}
