pub struct Entry { pub key: String, pub value: String }

fn parse_entry(e: &str) -> Option<Entry> {
    let parts: Vec<&str> = e.split('=').collect();
    if parts.len() != 2 || parts[0].is_empty() { return None; }
    Some(Entry { key: parts[0].to_string(), value: parts[1].to_string() })
}

pub fn parse_kv(s: &str) -> Option<Vec<Entry>> {
    if s.is_empty() { return Some(Vec::new()); }
    s.split(';').map(parse_entry).collect()
}
