package clock_test

import (
	"bytes"
	_ "embed"
	"fmt"
	. "github.com/Jumpaku/gkoku/clock"
	"testing"

	"github.com/Jumpaku/gkoku/internal/tests/assert"
)

func equalInstant(t *testing.T, want Instant, got Instant) {
	t.Helper()

	assert.Equal(t, want.State(), got.State())
	gs, gn := got.Unix()
	ws, wn := want.Unix()
	assert.Equal(t, ws, gs)
	assert.Equal(t, wn, gn)
}

//go:embed testcases/testdata/instant_add.txt
var testcasesInstantAdd []byte

func TestInstant_Add(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   Duration
		want Instant
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantAdd)
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
			name: fmt.Sprintf("Seconds(%d,%d).Add(Seconds(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Unix(sutSec, sutNano),
			in:   Seconds(inSec, inNano),
			want: Unix(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Add(testcase.in)

			equalInstant(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/instant_add_nano.txt
var testcasesInstantAddNano []byte

func TestInstant_AddNano(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   int64
		want Instant
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantAddNano)
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
			name: fmt.Sprintf("Seconds(%d,%d).AddNano(%d)", sutSec, sutNano, inNano),
			sut:  Unix(sutSec, sutNano),
			in:   inNano,
			want: Unix(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.AddNano(testcase.in)

			equalInstant(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/instant_sub.txt
var testcasesInstantSub []byte

func TestInstant_Sub(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   Duration
		want Instant
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantSub)
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
			name: fmt.Sprintf("Seconds(%d,%d).Sub(Seconds(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Unix(sutSec, sutNano),
			in:   Seconds(inSec, inNano),
			want: Unix(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Sub(testcase.in)

			equalInstant(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/instant_sub_nano.txt
var testcasesInstantSubNano []byte

func TestInstant_SubNano(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   int64
		want Instant
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantSubNano)
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
			name: fmt.Sprintf("Seconds(%d,%d).SubNano(%d)", sutSec, sutNano, inNano),
			sut:  Unix(sutSec, sutNano),
			in:   inNano,
			want: Unix(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.SubNano(testcase.in)

			equalInstant(t, testcase.want, got)
		})
	}
}

func TestInstant_Between(t *testing.T) {
	testcases := []struct {
		name   string
		sut    Instant
		inLow  Instant
		inHigh Instant
		want   bool
	}{
		{
			name:   "sut<h<l",
			sut:    Unix(10, 0),
			inLow:  Unix(30, 0),
			inHigh: Unix(20, 0),
			want:   false,
		},
		{
			name:   "h=sut<l",
			sut:    Unix(20, 0),
			inLow:  Unix(30, 0),
			inHigh: Unix(20, 0),
			want:   false,
		},
		{
			name:   "h<sut<l",
			sut:    Unix(25, 0),
			inLow:  Unix(30, 0),
			inHigh: Unix(20, 0),
			want:   false,
		},
		{
			name:   "h<sut=l",
			sut:    Unix(25, 0),
			inLow:  Unix(30, 0),
			inHigh: Unix(20, 0),
			want:   false,
		},
		{
			name:   "h<l<sut",
			sut:    Unix(40, 0),
			inLow:  Unix(30, 0),
			inHigh: Unix(20, 0),
			want:   false,
		},
		{
			name:   "sut<l=h",
			sut:    Unix(15, 0),
			inLow:  Unix(25, 0),
			inHigh: Unix(25, 0),
			want:   false,
		},
		{
			name:   "l=sut=h",
			sut:    Unix(25, 0),
			inLow:  Unix(25, 0),
			inHigh: Unix(25, 0),
			want:   true,
		},
		{
			name:   "l=h<sut",
			sut:    Unix(35, 0),
			inLow:  Unix(25, 0),
			inHigh: Unix(25, 0),
			want:   false,
		},
		{
			name:   "sut<l<h",
			sut:    Unix(10, 0),
			inLow:  Unix(20, 0),
			inHigh: Unix(30, 0),
			want:   false,
		},
		{
			name:   "l=sut<h",
			sut:    Unix(20, 0),
			inLow:  Unix(20, 0),
			inHigh: Unix(30, 0),
			want:   true,
		},
		{
			name:   "l<sut<h",
			sut:    Unix(25, 0),
			inLow:  Unix(20, 0),
			inHigh: Unix(30, 0),
			want:   true,
		},
		{
			name:   "l<sut=h",
			sut:    Unix(30, 0),
			inLow:  Unix(20, 0),
			inHigh: Unix(30, 0),
			want:   true,
		},
		{
			name:   "l<h<sut",
			sut:    Unix(40, 0),
			inLow:  Unix(20, 0),
			inHigh: Unix(30, 0),
			want:   false,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Between(testcase.inLow, testcase.inHigh)

			assert.Equal(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/instant_cmp.txt
var testcasesInstantCmp []byte

func TestInstant_Cmp(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   Instant
		want int
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantCmp)
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
			sut:  Unix(sutSec, sutNano),
			in:   Unix(inSec, inNano),
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

//go:embed testcases/testdata/instant_after.txt
var testcasesInstantAfter []byte

func TestInstant_After(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   Instant
		want bool
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantAfter)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, inSec, inNano int64
		var want bool
		if _, err := fmt.Fscanf(reader, "%d %d %d %d %t\n", &sutSec, &sutNano, &inSec, &inNano, &want); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Unix(%d,%d).After(Unix(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Unix(sutSec, sutNano),
			in:   Unix(inSec, inNano),
			want: want,
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.After(testcase.in)

			assert.Equal(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/instant_before.txt
var testcasesInstantBefore []byte

func TestInstant_Before(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   Instant
		want bool
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantBefore)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, inSec, inNano int64
		var want bool
		if _, err := fmt.Fscanf(reader, "%d %d %d %d %t\n", &sutSec, &sutNano, &inSec, &inNano, &want); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Unix(%d,%d).Before(Unix(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Unix(sutSec, sutNano),
			in:   Unix(inSec, inNano),
			want: want,
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Before(testcase.in)

			assert.Equal(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/instant_equal.txt
var testcasesInstantEqual []byte

func TestInstant_Equal(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   Instant
		want bool
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantEqual)
	if _, err := fmt.Fscanln(reader, &nTestcases); err != nil {
		t.Fatalf("failed to scan: %+v", err)
	}
	testcases := make([]testcase, nTestcases)
	for i := 0; i < nTestcases; i++ {
		var sutSec, sutNano, inSec, inNano int64
		var want bool
		if _, err := fmt.Fscanf(reader, "%d %d %d %d %t\n", &sutSec, &sutNano, &inSec, &inNano, &want); err != nil {
			t.Fatalf("failed to scan: %+v", err)
		}
		testcases[i] = testcase{
			name: fmt.Sprintf("Unix(%d,%d).Before(Unix(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Unix(sutSec, sutNano),
			in:   Unix(inSec, inNano),
			want: want,
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Equal(testcase.in)

			assert.Equal(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/instant_diff.txt
var testcasesInstantDiff []byte

func TestInstant_Diff(t *testing.T) {
	type testcase struct {
		name string
		sut  Instant
		in   Instant
		want Duration
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantDiff)
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
			name: fmt.Sprintf("Unix(%d,%d).Diff(Unix(%d,%d))", sutSec, sutNano, inSec, inNano),
			sut:  Unix(sutSec, sutNano),
			in:   Unix(inSec, inNano),
			want: Seconds(wantSec, WantNano),
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.sut.Diff(testcase.in)

			equalDuration(t, testcase.want, got)
		})
	}
}

//go:embed testcases/testdata/instant_unix.txt
var testcasesInstantUnix []byte

func TestInstant_Unix(t *testing.T) {
	type testcase struct {
		name        string
		sut         Instant
		wantSeconds int64
		wantNano    int64
	}
	var nTestcases int
	reader := bytes.NewBuffer(testcasesInstantUnix)
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
			name:        fmt.Sprintf("Unix(%d,%d).Unix()", sutSec, sutNano),
			sut:         Unix(sutSec, sutNano),
			wantSeconds: wantSeconds,
			wantNano:    wantNano,
		}
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			gotSeconds, gotNano := testcase.sut.Unix()

			assert.Equal(t, testcase.wantSeconds, gotSeconds)
			assert.Equal(t, testcase.wantNano, gotNano)
		})
	}
}

func TestInstant_State_OK(t *testing.T) {
	type testcase struct {
		name      string
		sut       Instant
		wantOK    bool
		wantState State
	}
	testcases := []testcase{
		{
			name:      "ok",
			sut:       Instant{},
			wantOK:    true,
			wantState: StateOK,
		},
		{
			name:      "ng",
			sut:       NGInstant(),
			wantOK:    false,
			wantState: StateOverflow,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			gotOK, gotState := testcase.sut.OK(), testcase.sut.State()

			assert.Equal(t, testcase.wantOK, gotOK)
			assert.Equal(t, testcase.wantState, gotState)
		})
	}
}
