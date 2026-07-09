fn main() {
    assert_eq!(rle(&[5,5,5,2]), vec![Run{v:5,n:3}, Run{v:2,n:1}]);
    assert_eq!(rle(&[]), vec![]);
    assert_eq!(rle(&[1,2,2]), vec![Run{v:1,n:1}, Run{v:2,n:2}]);
    println!("OK");
}
