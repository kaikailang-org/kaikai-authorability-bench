pub struct Point { pub x: i64, pub y: i64 }
pub struct Box { pub min_x: i64, pub min_y: i64, pub max_x: i64, pub max_y: i64 }

pub fn bounding_box(ps: &[Point]) -> Option<Box> {
    let first = ps.first()?;
    let mut b = Box { min_x: first.x, min_y: first.y, max_x: first.x, max_y: first.y };
    for p in ps {
        if p.x < b.min_x { b.min_x = p.x; }
        if p.y < b.min_y { b.min_y = p.y; }
        if p.x > b.max_x { b.max_x = p.x; }
        if p.y > b.max_y { b.max_y = p.y; }
    }
    Some(b)
}
