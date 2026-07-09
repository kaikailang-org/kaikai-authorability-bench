pub fn window_max_sum(xs: &[i64], k: usize) -> Option<i64> {
    if k == 0 || xs.len() < k { return None; }
    xs.windows(k).map(|w| w.iter().sum()).max()
}
