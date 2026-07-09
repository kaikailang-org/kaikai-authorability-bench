fn main() {
    assert_eq!(mask_digits("1234 5678 9012 3456"), "**** **** **** 3456");
    assert_eq!(mask_digits("12-34-5"), "*2-34-5");
    assert_eq!(mask_digits("12"), "12");
    assert_eq!(mask_digits(""), "");
    assert_eq!(mask_digits("abcd"), "abcd");
    assert_eq!(mask_digits("12345"), "*2345");
    println!("OK");
}
