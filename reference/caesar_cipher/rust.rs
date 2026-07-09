pub fn caesar(s: &str, k: i64) -> String {
    let r = ((k % 26) + 26) % 26;
    s.chars()
        .map(|c| {
            if c.is_ascii_lowercase() {
                (b'a' + ((c as u8 - b'a') as i64 + r) as u8 % 26) as char
            } else if c.is_ascii_uppercase() {
                (b'A' + ((c as u8 - b'A') as i64 + r) as u8 % 26) as char
            } else {
                c
            }
        })
        .collect()
}
