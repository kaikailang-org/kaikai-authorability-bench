pub fn longest_streak(days: &[bool]) -> i64 {
    let mut cur = 0i64;
    let mut best = 0i64;
    for &d in days {
        cur = if d { cur + 1 } else { 0 };
        if cur > best { best = cur; }
    }
    best
}
