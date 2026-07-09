func ParseDuration(s string) (int64, bool) {
	var num, total int64
	has, any := false, false
	last := int64(100000)
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int64(c-'0')
			has = true
			continue
		}
		var mult int64
		switch c {
		case 'h':
			mult = 3600
		case 'm':
			mult = 60
		case 's':
			mult = 1
		default:
			return 0, false
		}
		if !has || mult >= last {
			return 0, false
		}
		total += num * mult
		num, has, last, any = 0, false, mult, true
	}
	if has || !any {
		return 0, false
	}
	return total, true
}
