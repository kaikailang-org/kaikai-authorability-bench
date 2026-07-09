fn main() {
    assert_eq!(digital_root(0), 0);
    assert_eq!(digital_root(9), 9);
    assert_eq!(digital_root(942), 6);
    assert_eq!(digital_root(132189), 6);
    assert_eq!(digital_root(10), 1);
    assert_eq!(digital_root(999999999), 9);
    println!("OK");
}
