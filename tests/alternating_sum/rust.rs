fn main() {
    assert_eq!(alternating_sum(&[10,3,2]), 9);
    assert_eq!(alternating_sum(&[5,5]), 0);
    assert_eq!(alternating_sum(&[7]), 7);
    assert_eq!(alternating_sum(&[]), 0);
    assert_eq!(alternating_sum(&[1,2,3,4]), -2);
    assert_eq!(alternating_sum(&[-1,-2]), 1);
    println!("OK");
}
