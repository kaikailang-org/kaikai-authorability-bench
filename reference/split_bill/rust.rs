pub fn split_bill(total: i64, n: i64) -> Vec<i64> {
    let base = total / n;
    let r = total % n;
    (0..n).map(|i| base + if i < r { 1 } else { 0 }).collect()
}
