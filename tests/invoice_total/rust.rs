fn main() {
    assert_eq!(invoice_total(&[Line{qty:2,unit_cents:250}, Line{qty:1,unit_cents:100}]), 600);
    assert_eq!(invoice_total(&[]), 0);
    assert_eq!(invoice_total(&[Line{qty:0,unit_cents:999}]), 0);
    assert_eq!(invoice_total(&[Line{qty:-3,unit_cents:100}, Line{qty:4,unit_cents:25}]), 100);
    println!("OK");
}
