package prec

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

func equalDuration(t *testing.T, want Duration, got Duration) {
	t.Helper()

	assert.Equal(t, want.State(), got.State())
	gs, gn := got.Seconds()
	ws, wn := want.Seconds()
	assert.Equal(t, ws, gs)
	assert.Equal(t, wn, gn)
}

//go:embed testcases/testdata/duration_abs.txt
var testcasesDurationAbs []byte

func TestDuration_Abs(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		want Duration
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationAbs)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, wantSec, WantNano int64
		if _, err := fmt.Fscanf(reader, "%d %d %d %d\n", &sutSec, &sutNano, &wantSec, &WantNano); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Abs(Seconds(%d,%d))", sutSec, sutNano),
			sut:  Seconds(sutSec, sutNano),
			want: Seconds(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Abs()

			equalDuration(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/duration_sign.txt
var testcasesDurationSign []byte

func TestDuration_Sign(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		want int
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationSign)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano int64
		var want int
		if _, err := fmt.Fscanf(reader, "%d %d %d\n", &sutSec, &sutNano, &want); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Sign(Seconds(%d,%d))", sutSec, sutNano),
			sut:  Seconds(sutSec, sutNano),
			want: want,
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Sign()

			assert.Equal(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/duration_neg.txt
var testcasesDurationNag []byte

func TestDuration_Neg(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		want Duration
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationNag)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, wantSec, WantNano int64
		if _, err := fmt.Fscanf(reader, "%d %d %d %d\n", &sutSec, &sutNano, &wantSec, &WantNano); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Seconds(%d,%d).Neg()", sutSec, sutNano),
			sut:  Seconds(sutSec, sutNano),
			want: Seconds(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Neg()

			equalDuration(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/duration_add.txt
var testcasesDurationAdd []byte

func TestDuration_Add(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		in   Duration
		want Duration
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationAdd)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, inSec, inNano, wantSec, WantNano int64
		if _, err := fmt.Fscanf(reader, "%d %d %d %d %d %d\n", &sutSec, &sutNano, &inSec, &inNano, &wantSec, &WantNano); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Add(Seconds(%d,%d),Seconds(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Seconds(sutSec, sutNano),
			in:   Seconds(inSec, inNano),
			want: Seconds(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Add(testcase.in)

			equalDuration(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/duration_add_nano.txt
var testcasesDurationAddNano []byte

func TestDuration_AddNano(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		in   int64
		want Duration
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationAddNano)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, inNano, wantSec, WantNano int64
		if _, err := fmt.Fscanf(reader, "%d %d %d %d %d\n", &sutSec, &sutNano, &inNano, &wantSec, &WantNano); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("AddNano(Seconds(%d,%d),%d)", sutSec, sutNano, inNano),
			sut:  Seconds(sutSec, sutNano),
			in:   inNano,
			want: Seconds(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.AddNano(testcase.in)

			equalDuration(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/duration_sub.txt
var testcasesDurationSub []byte

func TestDuration_Sub(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		in   Duration
		want Duration
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationSub)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, inSec, inNano, wantSec, WantNano int64
		if _, err := fmt.Fscanf(reader, "%d %d %d %d %d %d\n", &sutSec, &sutNano, &inSec, &inNano, &wantSec, &WantNano); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Sub(Seconds(%d,%d),Seconds(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Seconds(sutSec, sutNano),
			in:   Seconds(inSec, inNano),
			want: Seconds(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Sub(testcase.in)

			equalDuration(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/duration_sub_nano.txt
var testcasesDurationSubNano []byte

func TestDuration_SubNano(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		in   int64
		want Duration
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationSubNano)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, inNano, wantSec, WantNano int64
		if _, err := fmt.Fscanf(reader, "%d %d %d %d %d\n", &sutSec, &sutNano, &inNano, &wantSec, &WantNano); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("SubNano(Seconds(%d,%d),%d)", sutSec, sutNano, inNano),
			sut:  Seconds(sutSec, sutNano),
			in:   inNano,
			want: Seconds(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.SubNano(testcase.in)

			equalDuration(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/duration_cmp.txt
var testcasesDurationCmp []byte

func TestDuration_Cmp(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		in   Duration
		want int
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationCmp)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, inSec, inNano int64
		var want int
		if _, err := fmt.Fscanf(reader, "%d %d %d %d %d\n", &sutSec, &sutNano, &inSec, &inNano, &want); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Seconds(%d,%d).Cmp(Seconds(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Seconds(sutSec, sutNano),
			in:   Seconds(inSec, inNano),
			want: want,
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Cmp(testcase.in)

			assert.Equal(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/duration_seconds.txt
var testcasesDurationSeconds []byte

func TestDuration_Seconds(t *testing.T) {
	type testcase struct {
		name        string
		sut         Duration
		wantSeconds int64
		wantNano    int64
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesDurationSeconds)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, wantSeconds, wantNano int64
		if _, err := fmt.Fscanf(reader, "%d %d %d %d\n", &sutSec, &sutNano, &wantSeconds, &wantNano); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name:        fmt.Sprintf("Seconds(%d,%d).Seconds()", sutSec, sutNano),
			sut:         Seconds(sutSec, sutNano),
			wantSeconds: wantSeconds,
			wantNano:    wantNano,
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			gotSeconds, gotNano := testcase.sut.Seconds()

			assert.Equal(t, testcase.wantSeconds, gotSeconds)
			assert.Equal(t, testcase.wantNano, gotNano)
		})
	}
}

func TestDuration_OK(t *testing.T) {
	type testcase struct {
		name string
		sut  Duration
		want bool
	}
	testcases := []testcase{
		{
			name: "ok",
			sut:  Duration{},
			want: true,
		},
		{
			name: "ng",
			sut:  Duration{state: StateOverflow},
			want: false,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.OK()

			assert.Equal(t, testcase.want, got)
		})
	}
}
