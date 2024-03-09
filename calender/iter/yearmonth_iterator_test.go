package iter

import (
	"github.com/Jumpaku/gkoku/calender"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewYearMonthIterator(t *testing.T) {
	type args struct {
		yearMonth calender.YearMonth
	}
	tests := []struct {
		name string
		args args
		want YearMonthIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, OfYearMonth(tt.args.yearMonth), "OfYearMonth(%v)", tt.args.yearMonth)
		})
	}
}

func Test_yearMonthIter_Date(t *testing.T) {
	type fields struct {
		yearMonth calender.YearMonth
	}
	type args struct {
		dayOfMonth int
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
			y := &yearMonthIter{
				yearMonth: tt.fields.yearMonth,
			}
			assert.Equalf(t, tt.want, y.Date(tt.args.dayOfMonth), "Date(%v)", tt.args.dayOfMonth)
		})
	}
}

func Test_yearMonthIter_Days(t *testing.T) {
	type fields struct {
		yearMonth calender.YearMonth
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
			y := &yearMonthIter{
				yearMonth: tt.fields.yearMonth,
			}
			assert.Equalf(t, tt.want, y.Days(), "Days()")
		})
	}
}

func Test_yearMonthIter_Diff(t *testing.T) {
	type fields struct {
		yearMonth calender.YearMonth
	}
	type args struct {
		from YearMonthIterator
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
			y := &yearMonthIter{
				yearMonth: tt.fields.yearMonth,
			}
			assert.Equalf(t, tt.want, y.Diff(tt.args.from), "Diff(%v)", tt.args.from)
		})
	}
}

func Test_yearMonthIter_FirstDate(t *testing.T) {
	type fields struct {
		yearMonth calender.YearMonth
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
			y := &yearMonthIter{
				yearMonth: tt.fields.yearMonth,
			}
			assert.Equalf(t, tt.want, y.FirstDate(), "FirstDate()")
		})
	}
}

func Test_yearMonthIter_Get(t *testing.T) {
	type fields struct {
		yearMonth calender.YearMonth
	}
	tests := []struct {
		name   string
		fields fields
		want   calender.YearMonth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yearMonthIter{
				yearMonth: tt.fields.yearMonth,
			}
			assert.Equalf(t, tt.want, y.Get(), "Get()")
		})
	}
}

func Test_yearMonthIter_LastDate(t *testing.T) {
	type fields struct {
		yearMonth calender.YearMonth
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
			y := &yearMonthIter{
				yearMonth: tt.fields.yearMonth,
			}
			assert.Equalf(t, tt.want, y.LastDate(), "LastDate()")
		})
	}
}

func Test_yearMonthIter_Move(t *testing.T) {
	type fields struct {
		yearMonth calender.YearMonth
	}
	type args struct {
		months int
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
			y := &yearMonthIter{
				yearMonth: tt.fields.yearMonth,
			}
			y.Move(tt.args.months)
		})
	}
}

func Test_yearMonthIter_Year(t *testing.T) {
	type fields struct {
		yearMonth calender.YearMonth
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
			y := &yearMonthIter{
				yearMonth: tt.fields.yearMonth,
			}
			assert.Equalf(t, tt.want, y.Year(), "Year()")
		})
	}
}
