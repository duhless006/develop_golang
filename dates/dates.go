package dates

const DaysInWeek int = 7

func WeekTodays(week int) int {
	return week * 7
}
func DaysToWeek(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
