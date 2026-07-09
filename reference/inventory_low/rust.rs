pub struct Stock { pub sku: i64, pub qty: i64 }

pub fn inventory_low(stock: &[Stock], threshold: i64) -> Vec<i64> {
    stock.iter().filter(|s| s.qty < threshold).map(|s| s.sku).collect()
}
