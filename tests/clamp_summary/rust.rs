fn main() {
    let a = clamp_summary(&[1, 50, -3, 7], 0, 10);
    assert_eq!(a.values, vec![1, 10, 0, 7]);
    assert_eq!(a.changed, 2);
    let b = clamp_summary(&[], 0, 10);
    assert_eq!(b.values, Vec::<i64>::new());
    assert_eq!(b.changed, 0);
    let c = clamp_summary(&[5, 5], 5, 5);
    assert_eq!(c.values, vec![5, 5]);
    assert_eq!(c.changed, 0);
    println!("OK");
}
