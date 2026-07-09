pub fn batch_withdraw(balance: i64, amounts: &[i64]) -> Option<i64> {
    let mut bal = balance;
    for &a in amounts {
        if a <= 0 || a > bal { return None; }
        bal -= a;
    }
    Some(bal)
}
