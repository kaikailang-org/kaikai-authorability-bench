fn main() {
    assert_eq!(temperature_swing(&[10,13,9]), 4);
    assert_eq!(temperature_swing(&[5]), 0);
    assert_eq!(temperature_swing(&[]), 0);
    assert_eq!(temperature_swing(&[-3,4]), 7);
    assert_eq!(temperature_swing(&[1,1,1]), 0);
    assert_eq!(temperature_swing(&[0,-10,5]), 15);
    println!("OK");
}
