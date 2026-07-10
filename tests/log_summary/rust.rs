fn is(s: &str, level: &str, msg: &str) -> bool {
    matches!(parse_line(s), Some(l) if l.level == level && l.message == msg)
}

fn main() {
    let mut ok = true;

    ok = ok && is("WARN: disk low", "WARN", "disk low");
    ok = ok && is("INFO: ", "INFO", "");
    ok = ok && is("ERROR: a: b: c", "ERROR", "a: b: c");

    ok = ok && parse_line("INFO:x").is_none();
    ok = ok && parse_line("DEBUG: x").is_none();
    ok = ok && parse_line("info: x").is_none();
    ok = ok && parse_line("INFO").is_none();
    ok = ok && parse_line("").is_none();
    ok = ok && parse_line(" INFO: x").is_none();

    let lines = ["INFO: a", "ERROR: b", "INFO: c", "WARN: "];
    ok = ok && count_level(&lines, "INFO") == Some(2);
    ok = ok && count_level(&[], "ERROR") == Some(0);
    ok = ok && count_level(&lines, "DEBUG").is_none();
    ok = ok && count_level(&["INFO: a", "broken"], "INFO").is_none();

    ok = ok && summary(&lines) == Some("INFO=2 WARN=1 ERROR=1".to_string());
    ok = ok && summary(&[]) == Some("INFO=0 WARN=0 ERROR=0".to_string());
    ok = ok && summary(&["INFO: a", "oops: b"]).is_none();

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
