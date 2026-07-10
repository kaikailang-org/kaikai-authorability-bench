fn add(book: &[Student], id: i64, score: i64) -> Vec<Student> {
    add_score(book, id, score).unwrap()
}

fn main() {
    let mut ok = true;

    let b = add(&[], 1, 90);
    let b = add(&b, 1, 95);
    ok = ok && b.len() == 1 && average_floor(&b, 1) == Some(92);

    let b = add(&b, 1, 92);
    ok = ok && average_floor(&b, 1) == Some(92);

    let b = add(&b, 2, 80);
    ok = ok && b.len() == 2 && b[0].id == 1 && b[1].id == 2;
    ok = ok && top_student(&b) == Some(1);

    let t = add(&add(&[], 1, 92), 2, 92);
    ok = ok && top_student(&t).is_none();
    ok = ok && top_student(&[]).is_none();

    let nb = vec![
        Student { id: 7, scores: vec![] },
        Student { id: 8, scores: vec![50] },
    ];
    ok = ok && top_student(&nb) == Some(8);
    ok = ok && average_floor(&nb, 7).is_none();
    ok = ok && average_floor(&nb, 99).is_none();

    ok = ok && add_score(&b, 1, 101).is_none();
    ok = ok && add_score(&b, 1, -1).is_none();
    let b2 = add(&b, 3, 0);
    ok = ok && average_floor(&b2, 3) == Some(0);
    let b3 = add(&b, 3, 100);
    ok = ok && average_floor(&b3, 3) == Some(100);

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
