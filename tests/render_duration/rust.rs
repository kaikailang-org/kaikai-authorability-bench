fn main() {
    assert_eq!(render_duration(3725), "1h02m05s");
    assert_eq!(render_duration(90), "1m30s");
    assert_eq!(render_duration(5), "5s");
    assert_eq!(render_duration(0), "0s");
    assert_eq!(render_duration(3600), "1h00m00s");
    assert_eq!(render_duration(60), "1m00s");
    assert_eq!(render_duration(45296), "12h34m56s");
    println!("OK");
}
