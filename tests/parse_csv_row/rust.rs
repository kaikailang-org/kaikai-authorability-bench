fn v(xs: &[&str]) -> Option<Vec<String>> {
    Some(xs.iter().map(|s| s.to_string()).collect())
}

fn main() {
    assert_eq!(parse_csv_row("a,b,c"), v(&["a", "b", "c"]));
    assert_eq!(parse_csv_row("\"a,b\",c"), v(&["a,b", "c"]));
    assert_eq!(parse_csv_row("a,\"x\"\"y\""), v(&["a", "x\"y"]));
    assert_eq!(parse_csv_row("\"\""), v(&[""]));
    assert_eq!(parse_csv_row(""), v(&[""]));
    assert_eq!(parse_csv_row("a,"), v(&["a", ""]));
    assert_eq!(parse_csv_row(",a"), v(&["", "a"]));
    assert_eq!(parse_csv_row(",,"), v(&["", "", ""]));
    assert_eq!(parse_csv_row("\"\"\"\""), v(&["\""]));
    assert_eq!(parse_csv_row("\"a\nb\",c"), v(&["a\nb", "c"]));
    assert_eq!(parse_csv_row("a,\"b\"x"), None);
    assert_eq!(parse_csv_row("a\"b"), None);
    assert_eq!(parse_csv_row("\"a"), None);
    assert_eq!(parse_csv_row("\"a\"\""), None);
    assert_eq!(parse_csv_row("x,\"a\"\""), None);
    println!("OK");
}
