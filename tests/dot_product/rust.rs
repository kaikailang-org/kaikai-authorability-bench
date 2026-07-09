fn main() {
    assert_eq!(dot_product(&[1,2,3], &[4,5,6]), Some(32));
    assert_eq!(dot_product(&[1], &[]), None);
    assert_eq!(dot_product(&[], &[]), Some(0));
    println!("OK");
}
