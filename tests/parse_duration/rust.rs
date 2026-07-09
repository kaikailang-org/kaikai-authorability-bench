fn main() {
    assert_eq!(parse_duration("2h30m15s"), Some(9015));
    assert_eq!(parse_duration("90s"), Some(90));
    assert_eq!(parse_duration("2h"), Some(7200));
    assert_eq!(parse_duration("1h0m5s"), Some(3605));
    assert_eq!(parse_duration("90m"), Some(5400));
    assert_eq!(parse_duration(""), None);
    assert_eq!(parse_duration("5x"), None);
    assert_eq!(parse_duration("m5"), None);
    assert_eq!(parse_duration("1s1h"), None);
    assert_eq!(parse_duration("2h30"), None);
    println!("OK");
}
