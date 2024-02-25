package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Equal[T any](t *testing.T, expected, actual T, msgAndArgs ...interface{}) bool {
	t.Helper()
	return assert.Equal(t, expected, actual, msgAndArgs...)
}
