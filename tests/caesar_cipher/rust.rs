fn main() {
    assert_eq!(caesar("abc", 1), "bcd");
    assert_eq!(caesar("xyz", 3), "abc");
    assert_eq!(caesar("abc", -1), "zab");
    assert_eq!(caesar("Hello, World!", 13), "Uryyb, Jbeyq!");
    assert_eq!(caesar("abc", 26), "abc");
    assert_eq!(caesar("", 5), "");
    println!("OK");
}
