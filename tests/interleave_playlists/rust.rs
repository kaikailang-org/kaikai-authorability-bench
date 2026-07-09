fn main() {
    assert_eq!(interleave(&[1,3,5], &[2,4]), vec![1,2,3,4,5]);
    assert_eq!(interleave(&[], &[7,8]), vec![7,8]);
    assert_eq!(interleave(&[1], &[]), vec![1]);
    assert_eq!(interleave(&[], &[]), Vec::<i64>::new());
    assert_eq!(interleave(&[1], &[7,8,9]), vec![1,7,8,9]);
    println!("OK");
}
