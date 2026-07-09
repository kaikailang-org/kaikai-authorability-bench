fn main() {
    assert_eq!(median_floor(&[3,1,2]), Some(2));
    assert_eq!(median_floor(&[4,1,3,2]), Some(2));
    assert_eq!(median_floor(&[]), None);
    assert_eq!(median_floor(&[5]), Some(5));
    assert_eq!(median_floor(&[-3,-4]), Some(-3));
    assert_eq!(median_floor(&[1,1,8,8]), Some(4));
    println!("OK");
}
