pub fn count_words(s: &str) -> i64 {
    let mut n = 0;
    let mut in_word = false;
    for c in s.chars() {
        let sp = c == ' ' || c == '\t';
        if !sp && !in_word { n += 1; }
        in_word = !sp;
    }
    n
}
