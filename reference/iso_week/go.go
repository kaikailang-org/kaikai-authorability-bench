type IsoWeek struct{ Year, Week int64 }

func daysFromCivil(y, m, d int64) int64 {
	if m <= 2 {
		y--
	}
	era := y / 400
	yoe := y - era*400
	var mp int64
	if m > 2 {
		mp = m - 3
	} else {
		mp = m + 9
	}
	doy := (153*mp+2)/5 + d - 1
	doe := yoe*365 + yoe/4 - yoe/100 + doy
	return era*146097 + doe - 719468
}

func yearOfDays(z int64) int64 {
	z += 719468
	era := z / 146097
	doe := z - era*146097
	yoe := (doe - doe/1460 + doe/36524 - doe/146096) / 365
	y := yoe + era*400
	doy := doe - (365*yoe + yoe/4 - yoe/100)
	mp := (5*doy + 2) / 153
	if mp >= 10 {
		return y + 1
	}
	return y
}

func IsoWeekOf(y, m, d int64) IsoWeek {
	days := daysFromCivil(y, m, d)
	dow := ((days+3)%7 + 7) % 7
	thu := days - dow + 3
	iy := yearOfDays(thu)
	week := (thu-daysFromCivil(iy, 1, 1))/7 + 1
	return IsoWeek{iy, week}
}
