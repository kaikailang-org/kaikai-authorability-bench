fn main() {
    assert_eq!(tile_size(12, 18), 6);
    assert_eq!(tile_size(7, 7), 7);
    assert_eq!(tile_size(13, 7), 1);
    assert_eq!(tile_size(100, 10), 10);
    assert_eq!(tile_size(1, 999), 1);
    assert_eq!(tile_size(18, 12), 6);
    println!("OK");
}
