pub fn max_subarray(xs: &[i64]) -> i64 {
    let mut best = 0; let mut cur = 0; let mut first = true;
    for &x in xs {
        if first { best = x; cur = x; first = false; }
        else { cur = if cur + x > x { cur + x } else { x }; if cur > best { best = cur; } }
    }
    best
}
