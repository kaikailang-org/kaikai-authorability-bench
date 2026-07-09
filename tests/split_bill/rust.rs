fn main() {
    assert_eq!(split_bill(100, 3), vec![34,33,33]);
    assert_eq!(split_bill(100, 4), vec![25,25,25,25]);
    assert_eq!(split_bill(5, 6), vec![1,1,1,1,1,0]);
    assert_eq!(split_bill(0, 2), vec![0,0]);
    assert_eq!(split_bill(7, 1), vec![7]);
    println!("OK");
}
