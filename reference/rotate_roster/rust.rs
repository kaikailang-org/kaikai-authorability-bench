pub fn rotate(xs: &[i64], k: usize) -> Vec<i64> {
    if xs.is_empty() { return Vec::new(); }
    let r = k % xs.len();
    xs[r..].iter().chain(xs[..r].iter()).copied().collect()
}
