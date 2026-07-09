pub fn alternating_sum(xs: &[i64]) -> i64 {
    xs.iter().enumerate().map(|(i, &x)| if i % 2 == 0 { x } else { -x }).sum()
}
