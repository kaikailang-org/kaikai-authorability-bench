pub fn int_sqrt(n: i64) -> i64 {
    let (mut lo, mut hi) = (0i64, n.min(1_000_000));
    while lo <= hi {
        let mid = (lo + hi) / 2;
        if mid * mid <= n { lo = mid + 1; } else { hi = mid - 1; }
    }
    hi
}
