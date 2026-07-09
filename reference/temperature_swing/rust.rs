pub fn temperature_swing(xs: &[i64]) -> i64 {
    xs.windows(2).map(|w| (w[0] - w[1]).abs()).max().unwrap_or(0)
}
