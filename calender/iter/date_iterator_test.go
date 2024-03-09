package iter

import (
	"github.com/Jumpaku/gkoku/calender"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dateIter_Diff(t *testing.T) {
	type fields struct {
		date calender.Date
	}
	type args struct {
		from DateIterator
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
			i := &dateIter{
				date: tt.fields.date,
			}
			assert.Equalf(t, tt.want, i.Diff(tt.args.from), "Diff(%v)", tt.args.from)
		})
	}
}

func Test_dateIter_Move(t *testing.T) {
	type fields struct {
		date calender.Date
	}
	type args struct {
		days int
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
			i := &dateIter{
				date: tt.fields.date,
			}
			i.Move(tt.args.days)
		})
	}
}

func Test_dateIter_Year(t *testing.T) {
	type fields struct {
		date calender.Date
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
			i := &dateIter{
				date: tt.fields.date,
			}
			assert.Equalf(t, tt.want, i.Year(), "Year()")
		})
	}
}

func Test_dateIter_YearMonth(t *testing.T) {
	type fields struct {
		date calender.Date
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
			i := &dateIter{
				date: tt.fields.date,
			}
			assert.Equalf(t, tt.want, i.YearMonth(), "YearMonth()")
		})
	}
}

func Test_dateIter_YearWeek(t *testing.T) {
	type fields struct {
		date calender.Date
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
			i := &dateIter{
				date: tt.fields.date,
			}
			assert.Equalf(t, tt.want, i.YearWeek(), "YearWeek()")
		})
	}
}
