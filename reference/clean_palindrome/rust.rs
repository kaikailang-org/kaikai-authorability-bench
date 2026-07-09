pub fn clean_palindrome(s: &str) -> bool {
    let cs: Vec<char> = s.chars()
        .filter(|c| c.is_ascii_alphanumeric())
        .map(|c| c.to_ascii_lowercase())
        .collect();
    cs.iter().eq(cs.iter().rev())
}
