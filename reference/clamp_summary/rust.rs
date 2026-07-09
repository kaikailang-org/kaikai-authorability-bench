pub struct Clamped { pub values: Vec<i64>, pub changed: i64 }

pub fn clamp_summary(xs: &[i64], lo: i64, hi: i64) -> Clamped {
    let values: Vec<i64> = xs.iter().map(|&x| x.clamp(lo, hi)).collect();
    let changed = xs.iter().filter(|&&x| x < lo || x > hi).count() as i64;
    Clamped { values, changed }
}
