fn at(t: &Turtle, x: i64, y: i64, d: i64) -> bool {
    t.x == x && t.y == y && t.dir == d
}

fn main() {
    let mut ok = true;

    ok = ok && at(&turtle_new(), 0, 0, 0);

    ok = ok && turn(&turtle_new(), -1).dir == 3;
    ok = ok && turn(&turtle_new(), 6).dir == 2;
    ok = ok && turn(&turtle_new(), -7).dir == 1;
    ok = ok && turn(&Turtle { x: 5, y: 5, dir: 3 }, 1).dir == 0;

    ok = ok && at(&forward(&turtle_new(), 3), 0, 3, 0);
    ok = ok && at(&forward(&Turtle { x: 0, y: 0, dir: 2 }, 2), 0, -2, 2);
    ok = ok && at(&forward(&Turtle { x: 1, y: 1, dir: 3 }, 4), -3, 1, 3);
    ok = ok && at(&forward(&Turtle { x: 9, y: 9, dir: 1 }, -5), 9, 9, 1);

    ok = ok && matches!(walk(&["F", "R", "F", "F"]), Some(t) if at(&t, 2, 1, 1));
    ok = ok && matches!(walk(&["L", "F", "F", "L", "F"]), Some(t) if at(&t, -2, -1, 2));
    ok = ok && matches!(walk(&[]), Some(t) if at(&t, 0, 0, 0));
    ok = ok && walk(&["F", "X"]).is_none();
    ok = ok && walk(&["f"]).is_none();

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
