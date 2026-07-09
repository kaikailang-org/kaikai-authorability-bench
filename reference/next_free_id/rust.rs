pub fn next_free_id(taken: &[i64]) -> i64 {
    let mut c = 0;
    while taken.contains(&c) { c += 1; }
    c
}
