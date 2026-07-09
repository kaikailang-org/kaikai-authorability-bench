fn main() {
    assert_eq!(parse_hex("0x1F"), Some(31));
    assert_eq!(parse_hex("0xff"), Some(255));
    assert_eq!(parse_hex("0x0"), Some(0));
    assert_eq!(parse_hex("0xDeadBeef"), Some(3735928559));
    assert_eq!(parse_hex("1F"), None);
    assert_eq!(parse_hex("0x"), None);
    assert_eq!(parse_hex("0xG1"), None);
    assert_eq!(parse_hex(""), None);
    println!("OK");
}
