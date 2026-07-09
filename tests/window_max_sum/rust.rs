fn main() {
    assert_eq!(window_max_sum(&[1,2,3,4], 2), Some(7));
    assert_eq!(window_max_sum(&[2,-1,2,3], 2), Some(5));
    assert_eq!(window_max_sum(&[1,2], 3), None);
    assert_eq!(window_max_sum(&[5], 1), Some(5));
    assert_eq!(window_max_sum(&[-1,-2,-3], 1), Some(-1));
    assert_eq!(window_max_sum(&[4,4,4], 3), Some(12));
    println!("OK");
}
