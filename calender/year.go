package calender

type Year int

var _ interface {
	IsLeap() bool
	Days() int
	ContainsDay(dayOfYear int) bool
	ContainsWeek(week int) bool
	Iter() YearIterator
} = Year(0)

func (y Year) IsLeap() bool {
	switch {
	case y%400 == 0:
		return true
	case y%100 == 0:
		return false
	case y%4 == 0:
		return true
	default:
		return false
	}
}
func (y Year) Days() int {
	if y.IsLeap() {
		return 366
	}
	return 365
}
func (y Year) Weeks() int {
	return int(YearWeekOf(int(y), 1).WeeksUntil(YearWeekOf(int(y+1), 1)))
}

func (y Year) Iter() YearIterator {
	return nil
}
func (y Year) ContainsDay(dayOfYear int) bool {
	return 1 <= dayOfYear && dayOfYear <= y.Days()
}

func (y Year) ContainsWeek(week int) bool {
	return 1 <= week && week <= y.Weeks()
}
