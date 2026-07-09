fn main() {
    let stock = [Stock{sku:1,qty:5}, Stock{sku:2,qty:0}, Stock{sku:3,qty:12}, Stock{sku:2,qty:1}];
    assert_eq!(inventory_low(&stock, 5), vec![2, 2]);
    assert_eq!(inventory_low(&stock, 6), vec![1, 2, 2]);
    assert_eq!(inventory_low(&stock, 0), Vec::<i64>::new());
    assert_eq!(inventory_low(&[], 10), Vec::<i64>::new());
    println!("OK");
}
