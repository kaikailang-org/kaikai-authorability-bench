pub fn tile_size(w: i64, h: i64) -> i64 {
    let (mut a, mut b) = (w, h);
    while b != 0 {
        let t = a % b;
        a = b;
        b = t;
    }
    a
}
