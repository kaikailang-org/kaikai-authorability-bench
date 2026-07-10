pub struct Meeting {
    pub start: i64,
    pub finish: i64,
}

pub fn min_rooms(ms: &[Meeting]) -> Option<i64> {
    if ms.iter().any(|m| m.start >= m.finish) {
        return None;
    }
    let mut best = 0i64;
    for m in ms {
        let c = ms
            .iter()
            .filter(|x| x.start <= m.start && m.start < x.finish)
            .count() as i64;
        if c > best {
            best = c;
        }
    }
    Some(best)
}
