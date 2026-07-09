fn main() {
    assert_eq!(count_words("  hello   world "), 2);
    assert_eq!(count_words(""), 0);
    assert_eq!(count_words("   "), 0);
    assert_eq!(count_words("one"), 1);
    assert_eq!(count_words("a\tb"), 2);
    assert_eq!(count_words("a b c d"), 4);
    println!("OK");
}
