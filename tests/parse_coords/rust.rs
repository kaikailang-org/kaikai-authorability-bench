fn chk(s: &str, wx: i64, wy: i64) -> bool {
    match parse_coords(s) { Some(p) => p.x == wx && p.y == wy, None => false }
}

fn bad(s: &str) -> bool { parse_coords(s).is_none() }

fn main() {
    assert!(chk("3,-4", 3, -4));
    assert!(chk("0,0", 0, 0));
    assert!(chk("-12,7", -12, 7));
    assert!(bad("3,"));
    assert!(bad("3, 4"));
    assert!(bad("3.5,1"));
    assert!(bad("--2,1"));
    assert!(bad("1,2,3"));
    assert!(bad(""));
    println!("OK");
}
