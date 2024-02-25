package testcases

//go:generate go run ./cmd/duration/abs/main.go testdata/duration_abs.txt
//go:generate go run ./cmd/duration/sign/main.go testdata/duration_sign.txt
//go:generate go run ./cmd/duration/neg/main.go testdata/duration_neg.txt
//go:generate go run ./cmd/duration/add/main.go testdata/duration_add.txt
//go:generate go run ./cmd/duration/add_nano/main.go testdata/duration_add_nano.txt
//go:generate go run ./cmd/duration/sub/main.go testdata/duration_sub.txt
//go:generate go run ./cmd/duration/sub_nano/main.go testdata/duration_sub_nano.txt
//go:generate go run ./cmd/duration/cmp/main.go testdata/duration_cmp.txt
//go:generate go run ./cmd/duration/seconds/main.go testdata/duration_seconds.txt
