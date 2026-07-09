fn main() {
    assert_eq!(title_case("hello WORLD"), "Hello World");
    assert_eq!(title_case("  a  b "), "  A  B ");
    assert_eq!(title_case("3rd place"), "3rd Place");
    assert_eq!(title_case(""), "");
    assert_eq!(title_case("mIxEd CaSe"), "Mixed Case");
    println!("OK");
}
