fn main() {
    assert_eq!(batch_withdraw(100, &[30,50]), Some(20));
    assert_eq!(batch_withdraw(100, &[30,80]), None);
    assert_eq!(batch_withdraw(100, &[]), Some(100));
    assert_eq!(batch_withdraw(100, &[100]), Some(0));
    assert_eq!(batch_withdraw(100, &[0]), None);
    assert_eq!(batch_withdraw(100, &[-5]), None);
    assert_eq!(batch_withdraw(0, &[1]), None);
    println!("OK");
}
