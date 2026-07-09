fn chk(total: i64, size: i64, wp: i64, wl: i64) -> bool {
    let p = paginate(total, size);
    p.pages == wp && p.last_page_items == wl
}

fn main() {
    assert!(chk(10, 3, 4, 1));
    assert!(chk(9, 3, 3, 3));
    assert!(chk(0, 5, 0, 0));
    assert!(chk(1, 1, 1, 1));
    assert!(chk(5, 10, 1, 5));
    println!("OK");
}
