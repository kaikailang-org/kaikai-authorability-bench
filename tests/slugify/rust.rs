fn main() {
    assert_eq!(slugify("Hello, World!"), "hello-world");
    assert_eq!(slugify("  42 is  GREAT "), "42-is-great");
    assert_eq!(slugify("---"), "");
    assert_eq!(slugify(""), "");
    assert_eq!(slugify("already-good"), "already-good");
    assert_eq!(slugify("A"), "a");
    println!("OK");
}
