fn main() {
    let mut ok = true;

    ok = ok && push(&[2, 1], 3) == vec![3, 2, 1];

    ok = ok && step(&[3, 2], "add") == Some(vec![5]);
    ok = ok && step(&[3, 2, 9], "mul") == Some(vec![6, 9]);
    ok = ok && step(&[1, 2], "swap") == Some(vec![2, 1]);
    ok = ok && step(&[7], "dup") == Some(vec![7, 7]);
    ok = ok && step(&[7, 8], "pop") == Some(vec![8]);
    ok = ok && step(&[1], "add").is_none();
    ok = ok && step(&[], "dup").is_none();
    ok = ok && step(&[1, 2], "sub").is_none();

    ok = ok && run(&["2", "3", "add"]) == Some(5);
    ok = ok && run(&["2", "3", "mul", "dup", "add"]) == Some(12);
    ok = ok && run(&["07"]) == Some(7);
    ok = ok && run(&["-4", "2", "add"]) == Some(-2);
    ok = ok && run(&["1", "2", "swap", "pop"]) == Some(2);
    ok = ok && run(&["4", "pop"]).is_none();
    ok = ok && run(&["swap"]).is_none();
    ok = ok && run(&["-"]).is_none();
    ok = ok && run(&[]).is_none();
    ok = ok && run(&["1", "2"]).is_none();
    ok = ok && run(&["+5"]).is_none();

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
