package errorhandling

import "testing"

func TestPanic(t *testing.T) {
	Panic(false)

	t.Logf("log after panic")
}
