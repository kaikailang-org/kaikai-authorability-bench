fn peq(a: &[Param], b: &[(&str, &str)]) -> bool {
    a.len() == b.len() && a.iter().zip(b).all(|(p, (k, v))| p.key == *k && p.value == *v)
}

fn main() {
    let mut ok = true;

    let p1 = parse_query("a=1&b=2").unwrap();
    ok = ok && peq(&p1, &[("a", "1"), ("b", "2")]);
    ok = ok && peq(&parse_query("k=a=b").unwrap(), &[("k", "a=b")]);
    let p3 = parse_query("a=").unwrap();
    ok = ok && peq(&p3, &[("a", "")]);
    ok = ok && peq(&parse_query("").unwrap(), &[]);

    ok = ok && parse_query("=x").is_none();
    ok = ok && parse_query("a=1&&b=2").is_none();
    ok = ok && parse_query("flag").is_none();
    ok = ok && parse_query("a=1&flag").is_none();

    ok = ok && get_param(&p1, "b") == Some("2".to_string());
    ok = ok && get_param(&p1, "z").is_none();
    let dup = parse_query("x=1&x=2").unwrap();
    ok = ok && get_param(&dup, "x") == Some("1".to_string());

    ok = ok && peq(&set_param(&p1, "a", "9"), &[("a", "9"), ("b", "2")]);
    ok = ok && peq(&set_param(&p1, "c", "3"), &[("a", "1"), ("b", "2"), ("c", "3")]);
    ok = ok && peq(&set_param(&dup, "x", "7"), &[("x", "7"), ("x", "2")]);
    ok = ok && peq(&p1, &[("a", "1"), ("b", "2")]);

    ok = ok && render_query(&p1) == "a=1&b=2";
    ok = ok && render_query(&[]) == "";
    ok = ok && render_query(&p3) == "a=";
    ok = ok && render_query(&parse_query("k=a=b&z=").unwrap()) == "k=a=b&z=";

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
