pub fn initials(name: &str) -> String {
    name.split(' ')
        .filter(|w| !w.is_empty())
        .filter_map(|w| w.chars().next())
        .map(|c| c.to_ascii_uppercase())
        .collect()
}
