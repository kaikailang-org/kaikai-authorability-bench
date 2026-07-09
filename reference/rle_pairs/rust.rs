#[derive(PartialEq, Debug)]
pub struct Run { pub v: i64, pub n: i64 }
pub fn rle(xs: &[i64]) -> Vec<Run> {
    let mut out: Vec<Run> = Vec::new();
    for &x in xs {
        let extend = matches!(out.last(), Some(r) if r.v == x);
        if extend { out.last_mut().unwrap().n += 1; }
        else { out.push(Run { v: x, n: 1 }); }
    }
    out
}
