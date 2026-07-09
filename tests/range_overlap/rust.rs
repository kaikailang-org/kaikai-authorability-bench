fn chk(a: (i64, i64), b: (i64, i64), want: (i64, i64)) -> bool {
    match range_overlap(&Range{start:a.0,finish:a.1}, &Range{start:b.0,finish:b.1}) {
        Some(r) => (r.start, r.finish) == want,
        None => false,
    }
}

fn none_of(a: (i64, i64), b: (i64, i64)) -> bool {
    range_overlap(&Range{start:a.0,finish:a.1}, &Range{start:b.0,finish:b.1}).is_none()
}

fn main() {
    assert!(chk((1,5), (3,8), (3,5)));
    assert!(none_of((1,3), (3,5)));
    assert!(none_of((1,2), (8,9)));
    assert!(chk((2,9), (3,4), (3,4)));
    assert!(chk((3,8), (1,5), (3,5)));
    println!("OK");
}
