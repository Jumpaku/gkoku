package gkoku_test

//go:generate go run ./internal/cmd/testcases/duration/abs/main.go testdata/duration_abs.txt
//go:generate go run ./internal/cmd/testcases/duration/sign/main.go testdata/duration_sign.txt
//go:generate go run ./internal/cmd/testcases/duration/neg/main.go testdata/duration_neg.txt
//go:generate go run ./internal/cmd/testcases/duration/add/main.go testdata/duration_add.txt
//go:generate go run ./internal/cmd/testcases/duration/add_nano/main.go testdata/duration_add_nano.txt
//go:generate go run ./internal/cmd/testcases/duration/sub/main.go testdata/duration_sub.txt
//go:generate go run ./internal/cmd/testcases/duration/sub_nano/main.go testdata/duration_sub_nano.txt
//go:generate go run ./internal/cmd/testcases/duration/cmp/main.go testdata/duration_cmp.txt
//go:generate go run ./internal/cmd/testcases/duration/seconds/main.go testdata/duration_seconds.txt

//go:generate go run ./internal/cmd/testcases/instant/unix/main.go testdata/instant_unix.txt
//go:generate go run ./internal/cmd/testcases/instant/add/main.go testdata/instant_add.txt
//go:generate go run ./internal/cmd/testcases/instant/add_nano/main.go testdata/instant_add_nano.txt
//go:generate go run ./internal/cmd/testcases/instant/sub/main.go testdata/instant_sub.txt
//go:generate go run ./internal/cmd/testcases/instant/sub_nano/main.go testdata/instant_sub_nano.txt
//go:generate go run ./internal/cmd/testcases/instant/diff/main.go testdata/instant_diff.txt
//go:generate go run ./internal/cmd/testcases/instant/cmp/main.go testdata/instant_cmp.txt
//go:generate go run ./internal/cmd/testcases/instant/equal/main.go testdata/instant_equal.txt
//go:generate go run ./internal/cmd/testcases/instant/after/main.go testdata/instant_after.txt
//go:generate go run ./internal/cmd/testcases/instant/before/main.go testdata/instant_before.txt
