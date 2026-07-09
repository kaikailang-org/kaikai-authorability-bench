fn main() {
    assert_eq!(next_free_id(&[0,1,3]), 2);
    assert_eq!(next_free_id(&[]), 0);
    assert_eq!(next_free_id(&[1,2,3]), 0);
    assert_eq!(next_free_id(&[0,0,1,1,2]), 3);
    assert_eq!(next_free_id(&[3,1,0,2]), 4);
    println!("OK");
}
