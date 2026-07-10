fn bal(l: &[Account], id: i64, want: i64) -> bool {
    balance_of(l, id) == Some(want)
}

fn main() {
    let mut ok = true;

    let l = open_account(&[], 1).unwrap();
    let l = open_account(&l, 2).unwrap();
    ok = ok && l.len() == 2 && l[0].id == 1 && l[1].id == 2;
    ok = ok && open_account(&l, 1).is_none();

    let l = deposit(&l, 1, 100).unwrap();
    ok = ok && bal(&l, 1, 100);

    let l = transfer(&l, 1, 2, 40).unwrap();
    ok = ok && bal(&l, 1, 60) && bal(&l, 2, 40);

    ok = ok && transfer(&l, 2, 2, 10).is_none();
    ok = ok && transfer(&l, 1, 9, 10).is_none();
    ok = ok && transfer(&l, 9, 1, 10).is_none();
    ok = ok && transfer(&l, 2, 1, 41).is_none();
    ok = ok && transfer(&l, 1, 2, 0).is_none();

    let l2 = withdraw(&l, 2, 40).unwrap();
    ok = ok && bal(&l2, 2, 0);
    ok = ok && withdraw(&l2, 2, 1).is_none();
    ok = ok && withdraw(&l2, 1, 0).is_none();
    ok = ok && withdraw(&l2, 3, 5).is_none();

    ok = ok && deposit(&l2, 1, -5).is_none();
    ok = ok && deposit(&l2, 42, 5).is_none();

    ok = ok && bal(&l, 1, 60) && bal(&l, 2, 40);
    ok = ok && balance_of(&l, 99).is_none();

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
