pub struct Run { pub v: i64, pub n: i64 }

pub fn rle_decode(runs: &[Run]) -> Vec<i64> {
    let mut out = Vec::new();
    for r in runs {
        for _ in 0..r.n { out.push(r.v); }
    }
    out
}
