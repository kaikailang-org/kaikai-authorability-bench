fn it(sku: i64, qty: i64, reserved: i64) -> Item {
    Item { sku, qty, reserved }
}

fn main() {
    let mut ok = true;

    let inv = add_stock(&[], 7, 10).unwrap();
    ok = ok && inv.len() == 1 && inv[0] == it(7, 10, 0);

    let inv = reserve(&inv, 7, 4).unwrap();
    ok = ok && inv[0] == it(7, 10, 4);

    let inv = ship(&inv, 7, 3).unwrap();
    ok = ok && inv[0] == it(7, 7, 1);
    ok = ok && available(&inv, 7) == Some(6);

    let inv2 = add_stock(&inv, 3, 5).unwrap();
    let inv2 = add_stock(&inv2, 7, 2).unwrap();
    ok = ok && inv2.len() == 2 && inv2[0] == it(7, 9, 1) && inv2[1] == it(3, 5, 0);

    let before = inv2.clone();
    ok = ok && reserve(&inv2, 7, 100).is_none() && inv2 == before;

    ok = ok && add_stock(&inv2, 1, 0).is_none();
    ok = ok && reserve(&inv2, 99, 1).is_none();
    ok = ok && reserve(&inv2, 7, 9).is_none();
    ok = ok && release(&inv2, 7, 2).is_none();
    ok = ok && ship(&inv2, 7, 2).is_none();
    ok = ok && available(&inv2, 42).is_none();

    let inv3 = add_stock(&[], 1, 2).unwrap();
    let inv3 = reserve(&inv3, 1, 2).unwrap();
    let inv3 = ship(&inv3, 1, 2).unwrap();
    ok = ok && inv3.len() == 1 && inv3[0] == it(1, 0, 0);

    let inv4 = add_stock(&[], 2, 3).unwrap();
    let inv4 = reserve(&inv4, 2, 3).unwrap();
    let inv4 = release(&inv4, 2, 3).unwrap();
    ok = ok && available(&inv4, 2) == Some(3);

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
