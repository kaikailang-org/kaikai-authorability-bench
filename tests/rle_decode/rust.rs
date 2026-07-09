fn main() {
    assert_eq!(rle_decode(&[Run{v:5,n:3}, Run{v:2,n:1}]), vec![5,5,5,2]);
    assert_eq!(rle_decode(&[]), Vec::<i64>::new());
    assert_eq!(rle_decode(&[Run{v:7,n:0}, Run{v:1,n:2}]), vec![1,1]);
    assert_eq!(rle_decode(&[Run{v:-3,n:2}]), vec![-3,-3]);
    println!("OK");
}
