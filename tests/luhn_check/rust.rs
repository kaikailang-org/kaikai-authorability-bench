fn main() {
    assert!(luhn_check(&[7,9,9,2,7,3,9,8,7,1,3]));
    assert!(!luhn_check(&[7,9,9,2,7,3,9,8,7,1,1]));
    assert!(!luhn_check(&[]));
    assert!(luhn_check(&[0]));
    println!("OK");
}
