fn main() {
    assert_eq!(running_min_balance(&[100,-30,-90,50]), -20);
    assert_eq!(running_min_balance(&[5,10]), 0);
    assert_eq!(running_min_balance(&[]), 0);
    assert_eq!(running_min_balance(&[-1,-1,-1]), -3);
    assert_eq!(running_min_balance(&[-5,10,-3]), -5);
    println!("OK");
}
