#[derive(Clone, PartialEq)]
pub struct Item {
    pub sku: i64,
    pub qty: i64,
    pub reserved: i64,
}

fn find(inv: &[Item], sku: i64) -> Option<usize> {
    inv.iter().position(|it| it.sku == sku)
}

pub fn add_stock(inv: &[Item], sku: i64, n: i64) -> Option<Vec<Item>> {
    if n < 1 {
        return None;
    }
    let mut out = inv.to_vec();
    match find(inv, sku) {
        Some(i) => out[i].qty += n,
        None => out.push(Item { sku, qty: n, reserved: 0 }),
    }
    Some(out)
}

pub fn reserve(inv: &[Item], sku: i64, n: i64) -> Option<Vec<Item>> {
    let i = find(inv, sku)?;
    if n < 1 || inv[i].qty - inv[i].reserved < n {
        return None;
    }
    let mut out = inv.to_vec();
    out[i].reserved += n;
    Some(out)
}

pub fn release(inv: &[Item], sku: i64, n: i64) -> Option<Vec<Item>> {
    let i = find(inv, sku)?;
    if n < 1 || inv[i].reserved < n {
        return None;
    }
    let mut out = inv.to_vec();
    out[i].reserved -= n;
    Some(out)
}

pub fn ship(inv: &[Item], sku: i64, n: i64) -> Option<Vec<Item>> {
    let i = find(inv, sku)?;
    if n < 1 || inv[i].reserved < n {
        return None;
    }
    let mut out = inv.to_vec();
    out[i].qty -= n;
    out[i].reserved -= n;
    Some(out)
}

pub fn available(inv: &[Item], sku: i64) -> Option<i64> {
    let i = find(inv, sku)?;
    Some(inv[i].qty - inv[i].reserved)
}
