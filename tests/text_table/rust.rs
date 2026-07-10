fn main() {
    let mut ok = true;

    let ws = col_widths(&["name", "qty"], &[vec!["ana", "3"], vec!["b", "12"]]);
    ok = ok && ws == Some(vec![4, 3]);

    let want = "name | qty\n---- | ---\nana  | 3\nb    | 12";
    ok = ok && render_table(&["name", "qty"], &[vec!["ana", "3"], vec!["b", "12"]])
        == Some(want.to_string());

    ok = ok && render_table(&["id"], &[]) == Some("id\n--".to_string());
    ok = ok && render_table(&["h"], &[vec!["wide"]]) == Some("h\n----\nwide".to_string());

    ok = ok && render_row(&["a", "b"], &[3, 5]) == "a   | b";
    ok = ok && render_row(&["x"], &[7]) == "x";

    ok = ok && col_widths(&[], &[]).is_none();
    ok = ok && col_widths(&["a"], &[vec!["x", "y"]]).is_none();
    ok = ok && render_table(&["a", "b"], &[vec!["only"]]).is_none();

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
