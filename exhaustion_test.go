package gronx

import (
	"testing"
	"time"
)

func TestExclusiveTickRejectsExhaustedYear(t *testing.T) {
	at := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)
	const expr = "0 0 1 1 * 2026"
	for _, test := range []struct {
		name   string
		search func(string, time.Time, bool) (time.Time, error)
	}{
		{"next", NextTickAfter},
		{"previous", PrevTickBefore},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.search(expr, at, false)
			if err == nil {
				t.Fatalf("exhausted year should return an error, got %s", got)
			}
			got, err = test.search(expr, at, true)
			if err != nil || !got.Equal(at) {
				t.Fatalf("inclusive lookup should retain current tick, got %s, %v", got, err)
			}
		})
	}
}
