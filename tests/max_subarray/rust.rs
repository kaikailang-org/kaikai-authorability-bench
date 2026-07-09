fn main() {
    assert_eq!(max_subarray(&[-2,1,-3,4,-1,2,1,-5,4]), 6);
    assert_eq!(max_subarray(&[]), 0);
    assert_eq!(max_subarray(&[-3,-1,-2]), -1);
    assert_eq!(max_subarray(&[5]), 5);
    println!("OK");
}
