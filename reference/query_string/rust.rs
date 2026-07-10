pub struct Param {
    pub key: String,
    pub value: String,
}

pub fn parse_query(s: &str) -> Option<Vec<Param>> {
    if s.is_empty() {
        return Some(Vec::new());
    }
    let mut out = Vec::new();
    for piece in s.split('&') {
        let (k, v) = piece.split_once('=')?;
        if k.is_empty() {
            return None;
        }
        out.push(Param { key: k.to_string(), value: v.to_string() });
    }
    Some(out)
}

pub fn get_param(ps: &[Param], k: &str) -> Option<String> {
    ps.iter().find(|p| p.key == k).map(|p| p.value.clone())
}

pub fn set_param(ps: &[Param], k: &str, v: &str) -> Vec<Param> {
    let mut out: Vec<Param> = ps
        .iter()
        .map(|p| Param { key: p.key.clone(), value: p.value.clone() })
        .collect();
    match out.iter_mut().find(|p| p.key == k) {
        Some(p) => p.value = v.to_string(),
        None => out.push(Param { key: k.to_string(), value: v.to_string() }),
    }
    out
}

pub fn render_query(ps: &[Param]) -> String {
    ps.iter()
        .map(|p| format!("{}={}", p.key, p.value))
        .collect::<Vec<_>>()
        .join("&")
}
