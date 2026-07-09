pub fn parse_hex(s: &str) -> Option<i64> {
    let rest = s.strip_prefix("0x")?;
    if rest.is_empty() { return None; }
    let mut acc: i64 = 0;
    for c in rest.chars() {
        let d = c.to_digit(16)? as i64;
        acc = acc * 16 + d;
    }
    Some(acc)
}
