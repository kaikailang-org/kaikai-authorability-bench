pub struct Item { pub sku: i64, pub cents: i64 }

pub fn lookup_price(items: &[Item], sku: i64) -> Option<i64> {
    items.iter().find(|i| i.sku == sku).map(|i| i.cents)
}
