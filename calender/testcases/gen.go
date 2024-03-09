package testcases

//go:generate go run ./cmd/date/yyyymmdd/main.go testdata/date_yyyymmdd.txt
//go:generate go run ./cmd/date/yyyywwd/main.go testdata/date_yyyywwd.txt
//go:generate go run ./cmd/date/yyyyddd/main.go testdata/date_yyyyddd.txt
//go:generate go run ./cmd/date/compare/main.go testdata/date_compare.txt
//go:generate go run ./cmd/date/until/main.go testdata/date_until.txt
//go:generate go run ./cmd/date/addsub/main.go testdata/date_addsub.txt

//go:generate go run ./cmd/yearmonth/yyyymm/main.go testdata/yearmonth_yyyymm.txt
//go:generate go run ./cmd/yearmonth/compare/main.go testdata/yearmonth_compare.txt
//go:generate go run ./cmd/yearmonth/day/main.go testdata/yearmonth_day.txt
//go:generate go run ./cmd/yearmonth/until/main.go testdata/yearmonth_until.txt
//go:generate go run ./cmd/yearmonth/addsub/main.go testdata/yearmonth_addsub.txt

//go:generate go run ./cmd/yearweek/yyyyww/main.go testdata/yearweek_yyyyww.txt
//go:generate go run ./cmd/yearweek/day/main.go testdata/yearweek_day.txt
//go:generate go run ./cmd/yearweek/compare/main.go testdata/yearweek_compare.txt

//go:generate go run ./cmd/date/conv/main.go testdata/date_conv.txt
//go:generate go run ./cmd/yearmonth/conv/main.go testdata/yearmonth_conv.txt
//go:generate go run ./cmd/yearweek/conv/main.go testdata/yearweek_conv.txt
