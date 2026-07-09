pub fn balanced_brackets(s: &str) -> bool {
    let mut stack: Vec<char> = Vec::new();
    for c in s.chars() {
        match c {
            '(' | '[' | '{' => stack.push(c),
            ')' | ']' | '}' => {
                let want = match c { ')' => '(', ']' => '[', _ => '{' };
                if stack.pop() != Some(want) { return false; }
            }
            _ => {}
        }
    }
    stack.is_empty()
}
