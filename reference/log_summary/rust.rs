pub struct LogLine {
    pub level: String,
    pub message: String,
}

pub fn parse_line(s: &str) -> Option<LogLine> {
    for lvl in ["INFO", "WARN", "ERROR"] {
        if let Some(rest) = s.strip_prefix(&format!("{lvl}: ")) {
            return Some(LogLine { level: lvl.to_string(), message: rest.to_string() });
        }
    }
    None
}

pub fn count_level(lines: &[&str], level: &str) -> Option<i64> {
    if !["INFO", "WARN", "ERROR"].contains(&level) {
        return None;
    }
    let mut n = 0;
    for s in lines {
        if parse_line(s)?.level == level {
            n += 1;
        }
    }
    Some(n)
}

pub fn summary(lines: &[&str]) -> Option<String> {
    let i = count_level(lines, "INFO")?;
    let w = count_level(lines, "WARN")?;
    let e = count_level(lines, "ERROR")?;
    Some(format!("INFO={i} WARN={w} ERROR={e}"))
}
