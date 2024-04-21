package date_test

//go:generate go run ./internal/cmd/testcases/date/yyyymmdd/main.go testdata/date_yyyymmdd.txt
//go:generate go run ./internal/cmd/testcases/date/yyyywwd/main.go testdata/date_yyyywwd.txt
//go:generate go run ./internal/cmd/testcases/date/yyyyddd/main.go testdata/date_yyyyddd.txt
//go:generate go run ./internal/cmd/testcases/date/compare/main.go testdata/date_compare.txt
//go:generate go run ./internal/cmd/testcases/date/until/main.go testdata/date_until.txt
//go:generate go run ./internal/cmd/testcases/date/addsub/main.go testdata/date_addsub.txt

//go:generate go run ./internal/cmd/testcases/yearmonth/yyyymm/main.go testdata/yearmonth_yyyymm.txt
//go:generate go run ./internal/cmd/testcases/yearmonth/compare/main.go testdata/yearmonth_compare.txt
//go:generate go run ./internal/cmd/testcases/yearmonth/day/main.go testdata/yearmonth_day.txt
//go:generate go run ./internal/cmd/testcases/yearmonth/until/main.go testdata/yearmonth_until.txt
//go:generate go run ./internal/cmd/testcases/yearmonth/addsub/main.go testdata/yearmonth_addsub.txt

//go:generate go run ./internal/cmd/testcases/yearweek/yyyyww/main.go testdata/yearweek_yyyyww.txt
//go:generate go run ./internal/cmd/testcases/yearweek/day/main.go testdata/yearweek_day.txt
//go:generate go run ./internal/cmd/testcases/yearweek/compare/main.go testdata/yearweek_compare.txt

//go:generate go run ./internal/cmd/testcases/date/conv/main.go testdata/date_conv.txt
//go:generate go run ./internal/cmd/testcases/yearmonth/conv/main.go testdata/yearmonth_conv.txt
//go:generate go run ./internal/cmd/testcases/yearweek/conv/main.go testdata/yearweek_conv.txt
