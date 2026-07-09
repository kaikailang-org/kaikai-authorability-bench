pub fn tank_level(cap: i64, deltas: &[i64]) -> i64 {
    let mut lvl = 0;
    for &d in deltas {
        lvl = (lvl + d).clamp(0, cap);
    }
    lvl
}
