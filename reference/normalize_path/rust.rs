pub fn normalize_path(s: &str) -> Option<String> {
    if !s.starts_with('/') {
        return None;
    }
    let mut stack: Vec<&str> = Vec::new();
    for c in s.split('/') {
        match c {
            "" | "." => {}
            ".." => {
                stack.pop();
            }
            name => stack.push(name),
        }
    }
    Some(format!("/{}", stack.join("/")))
}
