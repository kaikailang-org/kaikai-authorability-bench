pub struct Turtle {
    pub x: i64,
    pub y: i64,
    pub dir: i64,
}

pub fn turtle_new() -> Turtle {
    Turtle { x: 0, y: 0, dir: 0 }
}

pub fn turn(t: &Turtle, r: i64) -> Turtle {
    Turtle { x: t.x, y: t.y, dir: ((t.dir + r) % 4 + 4) % 4 }
}

pub fn forward(t: &Turtle, n: i64) -> Turtle {
    let n = n.max(0);
    match t.dir {
        0 => Turtle { x: t.x, y: t.y + n, dir: t.dir },
        1 => Turtle { x: t.x + n, y: t.y, dir: t.dir },
        2 => Turtle { x: t.x, y: t.y - n, dir: t.dir },
        _ => Turtle { x: t.x - n, y: t.y, dir: t.dir },
    }
}

pub fn walk(cmds: &[&str]) -> Option<Turtle> {
    let mut t = turtle_new();
    for c in cmds {
        t = match *c {
            "L" => turn(&t, -1),
            "R" => turn(&t, 1),
            "F" => forward(&t, 1),
            _ => return None,
        };
    }
    Some(t)
}
