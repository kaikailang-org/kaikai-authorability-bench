pub struct Range { pub start: i64, pub finish: i64 }

pub fn merge_ranges(rs: &[Range]) -> Vec<Range> {
    let mut s: Vec<Range> = rs.iter().map(|r| Range { start: r.start, finish: r.finish }).collect();
    s.sort_by_key(|r| r.start);
    let mut out: Vec<Range> = Vec::new();
    for r in s {
        match out.last_mut() {
            Some(last) if r.start <= last.finish => {
                if r.finish > last.finish { last.finish = r.finish; }
            }
            _ => out.push(r),
        }
    }
    out
}
