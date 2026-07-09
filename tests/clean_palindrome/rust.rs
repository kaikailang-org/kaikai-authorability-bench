fn main() {
    assert!(clean_palindrome("A man, a plan, a canal: Panama"));
    assert!(!clean_palindrome("race a car"));
    assert!(clean_palindrome(""));
    assert!(clean_palindrome("12 21"));
    assert!(clean_palindrome("No 'x' in Nixon"));
    assert!(!clean_palindrome("1a2"));
    println!("OK");
}
