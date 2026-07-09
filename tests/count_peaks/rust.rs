fn main() {
    assert_eq!(count_peaks(&[0,5,0,5,0]), 2);
    assert_eq!(count_peaks(&[1,3,2]), 1);
    assert_eq!(count_peaks(&[1,2,3]), 0);
    assert_eq!(count_peaks(&[]), 0);
    assert_eq!(count_peaks(&[1,2]), 0);
    assert_eq!(count_peaks(&[1,3,3,1]), 0);
    println!("OK");
}
