fn m(a: i64, b: i64) -> Meeting {
    Meeting { start: a, finish: b }
}

fn main() {
    assert_eq!(min_rooms(&[m(0, 30), m(5, 10), m(15, 20)]), Some(2));
    assert_eq!(min_rooms(&[m(1, 3), m(3, 5)]), Some(1));
    assert_eq!(min_rooms(&[m(7, 10), m(2, 4)]), Some(1));
    assert_eq!(min_rooms(&[m(1, 5), m(2, 5), m(3, 5)]), Some(3));
    assert_eq!(min_rooms(&[]), Some(0));
    assert_eq!(min_rooms(&[m(4, 9)]), Some(1));
    assert_eq!(min_rooms(&[m(1, 10), m(2, 3), m(4, 5), m(6, 7)]), Some(2));
    assert_eq!(min_rooms(&[m(0, 5), m(5, 10), m(5, 8)]), Some(2));
    assert_eq!(min_rooms(&[m(5, 5)]), None);
    assert_eq!(min_rooms(&[m(1, 3), m(6, 2)]), None);
    assert_eq!(min_rooms(&[m(-5, -1), m(-3, 2)]), Some(2));
    println!("OK");
}
