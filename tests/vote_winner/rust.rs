fn main() {
    assert_eq!(vote_winner(&[3,1,3,2,3]), Some(3));
    assert_eq!(vote_winner(&[1,2,1,2]), None);
    assert_eq!(vote_winner(&[]), None);
    assert_eq!(vote_winner(&[7]), Some(7));
    assert_eq!(vote_winner(&[5,5,4]), Some(5));
    println!("OK");
}
