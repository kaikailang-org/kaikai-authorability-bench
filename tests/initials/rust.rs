fn main() {
    assert_eq!(initials("grace brewster hopper"), "GBH");
    assert_eq!(initials(" ada   lovelace "), "AL");
    assert_eq!(initials(""), "");
    assert_eq!(initials("x"), "X");
    assert_eq!(initials("Alan Turing"), "AT");
    println!("OK");
}
