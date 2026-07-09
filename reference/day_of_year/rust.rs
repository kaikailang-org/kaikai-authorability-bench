fn is_leap(y: i64) -> bool { (y % 4 == 0 && y % 100 != 0) || y % 400 == 0 }

pub fn day_of_year(y: i64, m: i64, d: i64) -> i64 {
    let feb = if is_leap(y) { 29 } else { 28 };
    let lens = [31, feb, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    lens[..(m - 1) as usize].iter().sum::<i64>() + d
}
