pub fn luhn_check(digits: &[i64]) -> bool {
    if digits.is_empty() { return false; }
    let sum: i64 = digits.iter().rev().enumerate().map(|(i, &d)| {
        if i % 2 == 1 { let t = d * 2; if t > 9 { t - 9 } else { t } } else { d }
    }).sum();
    sum % 10 == 0
}
