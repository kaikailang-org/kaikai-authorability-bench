pub fn title_case(s: &str) -> String {
    let mut out = String::new();
    let mut in_word = false;
    for c in s.chars() {
        if c == ' ' {
            in_word = false;
            out.push(c);
        } else {
            out.push(if in_word { c.to_ascii_lowercase() } else { c.to_ascii_uppercase() });
            in_word = true;
        }
    }
    out
}
