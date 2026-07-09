fn main() {
    assert_eq!(collatz_steps(1), 0);
    assert_eq!(collatz_steps(2), 1);
    assert_eq!(collatz_steps(3), 7);
    assert_eq!(collatz_steps(27), 111);
    println!("OK");
}
