pub fn count_peaks(xs: &[i64]) -> i64 {
    let mut n = 0;
    for i in 1..xs.len().saturating_sub(1) {
        if xs[i] > xs[i-1] && xs[i] > xs[i+1] { n += 1; }
    }
    n
}
