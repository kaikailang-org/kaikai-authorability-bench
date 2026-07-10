pub fn push(st: &[i64], v: i64) -> Vec<i64> {
    let mut out = vec![v];
    out.extend_from_slice(st);
    out
}

fn lit(t: &str) -> Option<i64> {
    let s = t.strip_prefix('-').unwrap_or(t);
    if s.is_empty() || !s.bytes().all(|b| b.is_ascii_digit()) {
        return None;
    }
    t.parse::<i64>().ok()
}

pub fn step(st: &[i64], op: &str) -> Option<Vec<i64>> {
    match op {
        "add" | "mul" => {
            if st.len() < 2 {
                return None;
            }
            let (b, a) = (st[0], st[1]);
            let r = if op == "add" { a + b } else { a * b };
            let mut out = vec![r];
            out.extend_from_slice(&st[2..]);
            Some(out)
        }
        "swap" => {
            if st.len() < 2 {
                return None;
            }
            let mut out = vec![st[1], st[0]];
            out.extend_from_slice(&st[2..]);
            Some(out)
        }
        "dup" => {
            if st.is_empty() {
                return None;
            }
            let mut out = vec![st[0]];
            out.extend_from_slice(st);
            Some(out)
        }
        "pop" => {
            if st.is_empty() {
                return None;
            }
            Some(st[1..].to_vec())
        }
        _ => Some(push(st, lit(op)?)),
    }
}

pub fn run(prog: &[&str]) -> Option<i64> {
    let mut st: Vec<i64> = Vec::new();
    for op in prog {
        st = step(&st, op)?;
    }
    if st.len() == 1 {
        Some(st[0])
    } else {
        None
    }
}
