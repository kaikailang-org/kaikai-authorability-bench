fn chk(ps: &[Point], want: (i64, i64, i64, i64)) -> bool {
    match bounding_box(ps) {
        Some(b) => (b.min_x, b.min_y, b.max_x, b.max_y) == want,
        None => false,
    }
}

fn main() {
    assert!(chk(&[Point{x:1,y:5}, Point{x:-2,y:3}, Point{x:4,y:4}], (-2, 3, 4, 5)));
    assert!(chk(&[Point{x:7,y:-1}], (7, -1, 7, -1)));
    assert!(bounding_box(&[]).is_none());
    println!("OK");
}
