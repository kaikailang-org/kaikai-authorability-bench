fn main() {
    let items = [Item{sku:10,cents:250}, Item{sku:7,cents:100}, Item{sku:10,cents:999}];
    assert_eq!(lookup_price(&items, 7), Some(100));
    assert_eq!(lookup_price(&items, 10), Some(250));
    assert_eq!(lookup_price(&items, 3), None);
    assert_eq!(lookup_price(&[], 1), None);
    println!("OK");
}
