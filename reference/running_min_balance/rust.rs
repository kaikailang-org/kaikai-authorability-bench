pub fn running_min_balance(txs: &[i64]) -> i64 {
    let mut bal = 0;
    let mut low = 0;
    for &t in txs {
        bal += t;
        if bal < low { low = bal; }
    }
    low
}
