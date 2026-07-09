fn main() {
    assert_eq!(safe_average(&[1,2,4]), Some(2));
    assert_eq!(safe_average(&[]), None);
    assert_eq!(safe_average(&[5]), Some(5));
    assert_eq!(safe_average(&[-7,0]), Some(-3));
    assert_eq!(safe_average(&[2,2,2,2]), Some(2));
    println!("OK");
}
