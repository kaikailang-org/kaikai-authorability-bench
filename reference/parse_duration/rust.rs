pub fn parse_duration(s: &str) -> Option<i64> {
    let mut num: i64 = 0;
    let mut has = false;
    let mut last: i64 = 100_000;
    let mut total: i64 = 0;
    let mut any = false;
    for c in s.chars() {
        if let Some(d) = c.to_digit(10) {
            num = num * 10 + d as i64;
            has = true;
        } else {
            let mult = match c { 'h' => 3600, 'm' => 60, 's' => 1, _ => return None };
            if !has || mult >= last { return None; }
            total += num * mult;
            num = 0; has = false; last = mult; any = true;
        }
    }
    if has || !any { None } else { Some(total) }
}
