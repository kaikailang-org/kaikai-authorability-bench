fn main() {
    assert_eq!(rotate(&[1,2,3,4], 1), vec![2,3,4,1]);
    assert_eq!(rotate(&[1,2,3,4], 0), vec![1,2,3,4]);
    assert_eq!(rotate(&[1,2,3,4], 4), vec![1,2,3,4]);
    assert_eq!(rotate(&[1,2,3,4], 6), vec![3,4,1,2]);
    assert_eq!(rotate(&[], 3), Vec::<i64>::new());
    println!("OK");
}
