fn main() {
    assert_eq!(int_sqrt(0), 0);
    assert_eq!(int_sqrt(1), 1);
    assert_eq!(int_sqrt(8), 2);
    assert_eq!(int_sqrt(9), 3);
    assert_eq!(int_sqrt(999_999_999_999), 999_999);
    assert_eq!(int_sqrt(1_000_000_000_000), 1_000_000);
    assert_eq!(int_sqrt(2), 1);
    println!("OK");
}
