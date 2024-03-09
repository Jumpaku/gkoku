package iter

import (
	"github.com/Jumpaku/gkoku/calender"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewYearWeekIterator(t *testing.T) {
	type args struct {
		yearWeek calender.YearWeek
	}
	tests := []struct {
		name string
		args args
		want YearWeekIterator
	}{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, OfYearWeek(tt.args.yearWeek), "OfYearWeek(%v)", tt.args.yearWeek)
		})
	}
}

func Test_yearWeekIter_Date(t *testing.T) {
	type fields struct {
		yearWeek calender.YearWeek
	}
	type args struct {
		dayOfWeek calender.DayOfWeek
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
			y := &yearWeekIter{
				yearWeek: tt.fields.yearWeek,
			}
			assert.Equalf(t, tt.want, y.Date(tt.args.dayOfWeek), "Date(%v)", tt.args.dayOfWeek)
		})
	}
}

func Test_yearWeekIter_Diff(t *testing.T) {
	type fields struct {
		yearWeek calender.YearWeek
	}
	type args struct {
		from YearWeekIterator
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
			y := &yearWeekIter{
				yearWeek: tt.fields.yearWeek,
			}
			assert.Equalf(t, tt.want, y.Diff(tt.args.from), "Diff(%v)", tt.args.from)
		})
	}
}

func Test_yearWeekIter_FirstDate(t *testing.T) {
	type fields struct {
		yearWeek calender.YearWeek
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
			y := &yearWeekIter{
				yearWeek: tt.fields.yearWeek,
			}
			assert.Equalf(t, tt.want, y.FirstDate(), "FirstDate()")
		})
	}
}

func Test_yearWeekIter_Get(t *testing.T) {
	type fields struct {
		yearWeek calender.YearWeek
	}
	tests := []struct {
		name   string
		fields fields
		want   calender.YearWeek
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearWeekIter{
				yearWeek: tt.fields.yearWeek,
			}
			assert.Equalf(t, tt.want, y.Get(), "Get()")
		})
	}
}

func Test_yearWeekIter_LastDate(t *testing.T) {
	type fields struct {
		yearWeek calender.YearWeek
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
			y := &yearWeekIter{
				yearWeek: tt.fields.yearWeek,
			}
			assert.Equalf(t, tt.want, y.LastDate(), "LastDate()")
		})
	}
}

func Test_yearWeekIter_Move(t *testing.T) {
	type fields struct {
		yearWeek calender.YearWeek
	}
	type args struct {
		weeks int
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
			y := &yearWeekIter{
				yearWeek: tt.fields.yearWeek,
			}
			y.Move(tt.args.weeks)
		})
	}
}

func Test_yearWeekIter_Year(t *testing.T) {
	type fields struct {
		yearWeek calender.YearWeek
	}
	tests := []struct {
		name   string
		fields fields
		want   YearIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearWeekIter{
				yearWeek: tt.fields.yearWeek,
			}
			assert.Equalf(t, tt.want, y.Year(), "Year()")
		})
	}
}
