pub struct IsoWeek {
    pub year: i64,
    pub week: i64,
}

fn days_from_civil(mut y: i64, m: i64, d: i64) -> i64 {
    if m <= 2 {
        y -= 1;
    }
    let era = y / 400;
    let yoe = y - era * 400;
    let mp = if m > 2 { m - 3 } else { m + 9 };
    let doy = (153 * mp + 2) / 5 + d - 1;
    let doe = yoe * 365 + yoe / 4 - yoe / 100 + doy;
    era * 146097 + doe - 719468
}

fn year_of_days(z: i64) -> i64 {
    let z = z + 719468;
    let era = z / 146097;
    let doe = z - era * 146097;
    let yoe = (doe - doe / 1460 + doe / 36524 - doe / 146096) / 365;
    let y = yoe + era * 400;
    let doy = doe - (365 * yoe + yoe / 4 - yoe / 100);
    let mp = (5 * doy + 2) / 153;
    if mp >= 10 {
        y + 1
    } else {
        y
    }
}

pub fn iso_week(y: i64, m: i64, d: i64) -> IsoWeek {
    let days = days_from_civil(y, m, d);
    let dow = ((days + 3) % 7 + 7) % 7;
    let thu = days - dow + 3;
    let iy = year_of_days(thu);
    let week = (thu - days_from_civil(iy, 1, 1)) / 7 + 1;
    IsoWeek { year: iy, week }
}
