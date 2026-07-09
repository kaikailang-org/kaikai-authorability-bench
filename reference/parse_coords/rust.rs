pub struct Point { pub x: i64, pub y: i64 }

fn int_val(s: &str) -> Option<i64> {
    let (neg, digits) = match s.strip_prefix('-') {
        Some(rest) => (true, rest),
        None => (false, s),
    };
    if digits.is_empty() || !digits.chars().all(|c| c.is_ascii_digit()) { return None; }
    let mut v: i64 = 0;
    for c in digits.chars() { v = v * 10 + (c as i64 - '0' as i64); }
    Some(if neg { -v } else { v })
}

pub fn parse_coords(s: &str) -> Option<Point> {
    let parts: Vec<&str> = s.split(',').collect();
    if parts.len() != 2 { return None; }
    Some(Point { x: int_val(parts[0])?, y: int_val(parts[1])? })
}
