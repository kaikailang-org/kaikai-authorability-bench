fn main() {
    assert_eq!(histogram_buckets(&[5,15,25,99,0,9]), vec![3,1,1,0,0,0,0,0,0,1]);
    assert_eq!(histogram_buckets(&[]), vec![0;10]);
    assert_eq!(histogram_buckets(&[10,19,90]), vec![0,2,0,0,0,0,0,0,0,1]);
    assert_eq!(histogram_buckets(&[0]), vec![1,0,0,0,0,0,0,0,0,0]);
    println!("OK");
}
