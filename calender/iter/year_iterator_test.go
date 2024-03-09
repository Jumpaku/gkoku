package iter

import (
	"github.com/Jumpaku/gkoku/calender"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewYearIterator(t *testing.T) {
	type args struct {
		year calender.Year
	}
	tests := []struct {
		name string
		args args
		want YearIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, OfYear(tt.args.year), "OfYear(%v)", tt.args.year)
		})
	}
}

func Test_yearIter_Date(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	type args struct {
		dayOfYear int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   DateIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.Date(tt.args.dayOfYear), "Date(%v)", tt.args.dayOfYear)
		})
	}
}

func Test_yearIter_Days(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.Days(), "Days()")
		})
	}
}

func Test_yearIter_Diff(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	type args struct {
		from YearIterator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.Diff(tt.args.from), "Diff(%v)", tt.args.from)
		})
	}
}

func Test_yearIter_FirstDate(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   DateIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.FirstDate(), "FirstDate()")
		})
	}
}

func Test_yearIter_FirstMonth(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   YearMonthIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.FirstMonth(), "FirstMonth()")
		})
	}
}

func Test_yearIter_FirstWeek(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   YearWeekIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.FirstWeek(), "FirstWeek()")
		})
	}
}

func Test_yearIter_Get(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   calender.Year
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.Get(), "Get()")
		})
	}
}

func Test_yearIter_LastDate(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   DateIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.LastDate(), "LastDate()")
		})
	}
}

func Test_yearIter_LastMonth(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   YearMonthIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.LastMonth(), "LastMonth()")
		})
	}
}

func Test_yearIter_LastWeek(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   YearWeekIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.LastWeek(), "LastWeek()")
		})
	}
}

func Test_yearIter_Month(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	type args struct {
		month calender.Month
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   YearMonthIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.Month(tt.args.month), "Month(%v)", tt.args.month)
		})
	}
}

func Test_yearIter_Move(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	type args struct {
		years int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			y.Move(tt.args.years)
		})
	}
}

func Test_yearIter_Week(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	type args struct {
		week int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   YearWeekIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.Week(tt.args.week), "Week(%v)", tt.args.week)
		})
	}
}

func Test_yearIter_Weeks(t *testing.T) {
	type fields struct {
		year calender.Year
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearIter{
				year: tt.fields.year,
			}
			assert.Equalf(t, tt.want, y.Weeks(), "Weeks()")
		})
	}
}
