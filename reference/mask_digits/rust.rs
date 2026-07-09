pub fn mask_digits(s: &str) -> String {
    let nd = s.chars().filter(|c| c.is_ascii_digit()).count();
    let mut to_mask = nd.saturating_sub(4);
    s.chars()
        .map(|c| {
            if c.is_ascii_digit() && to_mask > 0 {
                to_mask -= 1;
                '*'
            } else {
                c
            }
        })
        .collect()
}
