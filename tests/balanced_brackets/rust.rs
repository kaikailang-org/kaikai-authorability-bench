fn main() {
    assert!(balanced_brackets("([]{})"));
    assert!(balanced_brackets("f(a[0]) { return; }"));
    assert!(!balanced_brackets("([)]"));
    assert!(!balanced_brackets("("));
    assert!(balanced_brackets(""));
    assert!(!balanced_brackets(")("));
    assert!(balanced_brackets("no brackets at all"));
    println!("OK");
}
