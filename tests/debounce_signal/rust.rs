fn main() {
    assert_eq!(debounce(&[4,4,7,7,7,4]), vec![4,7,4]);
    assert_eq!(debounce(&[]), Vec::<i64>::new());
    assert_eq!(debounce(&[9]), vec![9]);
    assert_eq!(debounce(&[2,2,2,2]), vec![2]);
    assert_eq!(debounce(&[1,2,3]), vec![1,2,3]);
    println!("OK");
}
