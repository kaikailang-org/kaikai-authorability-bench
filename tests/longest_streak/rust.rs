fn main() {
    assert_eq!(longest_streak(&[true,true,false,true,true,true]), 3);
    assert_eq!(longest_streak(&[]), 0);
    assert_eq!(longest_streak(&[false,false]), 0);
    assert_eq!(longest_streak(&[true,true,true,true]), 4);
    assert_eq!(longest_streak(&[true,false,true]), 1);
    println!("OK");
}
