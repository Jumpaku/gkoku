package tests

import (
	"fmt"
	"io"
)

type Scanner struct {
	Data io.Reader
}

func (s *Scanner) ScanInt() int {
	var read int
	_, err := fmt.Fscanln(s.Data, &read)
	if err != nil {
		panic(fmt.Sprintf(`failed to read line: %+v`, err))
	}
	return read
}

func (s *Scanner) ScanInts(n int) []int {
	read := make([]int, n)
	ptr := []any{}
	for i := range read {
		ptr = append(ptr, &read[i])
	}
	_, err := fmt.Fscanln(s.Data, ptr...)
	if err != nil {
		panic(fmt.Sprintf(`failed to read line: %+v`, err))
	}
	return read
}

func (s *Scanner) ScanInt64s(n int) []int64 {
	read := make([]int64, n)
	ptr := []any{}
	for i := range read {
		ptr = append(ptr, &read[i])
	}
	_, err := fmt.Fscanln(s.Data, ptr...)
	if err != nil {
		panic(fmt.Sprintf(`failed to read line: %+v`, err))
	}
	return read
}
