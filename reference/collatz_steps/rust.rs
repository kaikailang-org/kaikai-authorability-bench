pub fn collatz_steps(n: i64) -> i64 {
    let mut n = n; let mut s = 0;
    while n != 1 { n = if n % 2 == 0 { n / 2 } else { 3 * n + 1 }; s += 1; }
    s
}
