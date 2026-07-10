fn lit(t: &str) -> Option<i64> {
    let s = t.strip_prefix('-').unwrap_or(t);
    if s.is_empty() || !s.bytes().all(|b| b.is_ascii_digit()) {
        return None;
    }
    t.parse::<i64>().ok()
}

pub fn eval_rpn(tokens: &[&str]) -> Option<i64> {
    let mut st: Vec<i64> = Vec::new();
    for t in tokens {
        match *t {
            "+" | "-" | "*" | "/" => {
                let b = st.pop()?;
                let a = st.pop()?;
                let r = match *t {
                    "+" => a + b,
                    "-" => a - b,
                    "*" => a * b,
                    _ => {
                        if b == 0 {
                            return None;
                        }
                        a / b
                    }
                };
                st.push(r);
            }
            _ => st.push(lit(t)?),
        }
    }
    if st.len() == 1 {
        Some(st[0])
    } else {
        None
    }
}
