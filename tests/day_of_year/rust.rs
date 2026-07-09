fn main() {
    assert_eq!(day_of_year(2024, 3, 1), 61);
    assert_eq!(day_of_year(2023, 3, 1), 60);
    assert_eq!(day_of_year(2000, 1, 1), 1);
    assert_eq!(day_of_year(1900, 3, 1), 60);
    assert_eq!(day_of_year(2024, 12, 31), 366);
    assert_eq!(day_of_year(2023, 12, 31), 365);
    println!("OK");
}
