pub struct Range { pub start: i64, pub finish: i64 }

pub fn range_overlap(a: &Range, b: &Range) -> Option<Range> {
    let lo = a.start.max(b.start);
    let hi = a.finish.min(b.finish);
    if lo < hi { Some(Range { start: lo, finish: hi }) } else { None }
}
