fn good(s: &str, want: &[(&str, &str)]) -> bool {
    match parse_kv(s) {
        Some(es) => es.len() == want.len()
            && es.iter().zip(want).all(|(e, w)| e.key == w.0 && e.value == w.1),
        None => false,
    }
}

fn bad(s: &str) -> bool { parse_kv(s).is_none() }

fn main() {
    assert!(good("a=1;b=2", &[("a", "1"), ("b", "2")]));
    assert!(good("a=", &[("a", "")]));
    assert!(good("", &[]));
    assert!(good("host=localhost", &[("host", "localhost")]));
    assert!(bad("=1"));
    assert!(bad("a"));
    assert!(bad("a=1;"));
    assert!(bad("a=b=c"));
    println!("OK");
}
