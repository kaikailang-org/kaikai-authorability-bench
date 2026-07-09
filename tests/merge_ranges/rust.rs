fn ranges_eq(a: &[Range], b: &[(i64, i64)]) -> bool {
    a.len() == b.len() && a.iter().zip(b).all(|(r, w)| r.start == w.0 && r.finish == w.1)
}

fn main() {
    assert!(ranges_eq(&merge_ranges(&[Range{start:5,finish:7}, Range{start:1,finish:3}, Range{start:2,finish:4}]),
        &[(1,4),(5,7)]));
    assert!(ranges_eq(&merge_ranges(&[Range{start:1,finish:2}, Range{start:2,finish:3}]), &[(1,3)]));
    assert!(ranges_eq(&merge_ranges(&[]), &[]));
    assert!(ranges_eq(&merge_ranges(&[Range{start:1,finish:10}, Range{start:2,finish:3}]), &[(1,10)]));
    assert!(ranges_eq(&merge_ranges(&[Range{start:4,finish:5}, Range{start:1,finish:2}]), &[(1,2),(4,5)]));
    println!("OK");
}
