func isLeap(y int64) bool { return (y%4 == 0 && y%100 != 0) || y%400 == 0 }

func DayOfYear(y, m, d int64) int64 {
	feb := int64(28)
	if isLeap(y) {
		feb = 29
	}
	lens := []int64{31, feb, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	total := d
	for i := int64(0); i < m-1; i++ {
		total += lens[i]
	}
	return total
}
