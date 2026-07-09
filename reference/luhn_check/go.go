func LuhnCheck(digits []int64) bool {
	if len(digits) == 0 { return false }
	var sum int64
	for i := 0; i < len(digits); i++ {
		d := digits[len(digits)-1-i]
		if i%2 == 1 {
			d *= 2
			if d > 9 { d -= 9 }
		}
		sum += d
	}
	return sum%10 == 0
}
