fn main() {
    assert_eq!(tank_level(10, &[7,7,-20,3]), 3);
    assert_eq!(tank_level(10, &[]), 0);
    assert_eq!(tank_level(0, &[5,-2]), 0);
    assert_eq!(tank_level(100, &[30,30]), 60);
    assert_eq!(tank_level(5, &[9]), 5);
    assert_eq!(tank_level(5, &[-3]), 0);
    println!("OK");
}
