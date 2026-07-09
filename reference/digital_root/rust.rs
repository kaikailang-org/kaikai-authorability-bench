pub fn digital_root(n: i64) -> i64 {
    let mut n = n;
    while n >= 10 {
        let mut s = 0;
        let mut m = n;
        while m > 0 { s += m % 10; m /= 10; }
        n = s;
    }
    n
}
