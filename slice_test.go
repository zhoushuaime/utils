package utils

import (
	"math"
	"testing"
)

func TestFindFirst(t *testing.T) {
	tests := []struct {
		name      string
		in        []string
		want      string
		wantIndex int
	}{
		{name: "duplicate matches", in: []string{"x", "target", "target"}, want: "target", wantIndex: 1},
		{name: "missing target", in: []string{"x", "y"}, wantIndex: -1},
		{name: "nil slice", wantIndex: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, index := FindFirst(tt.in, "target")
			if got != tt.want || index != tt.wantIndex {
				t.Fatalf("FindFirst(%v, %q) = (%q, %d), want (%q, %d)", tt.in, "target", got, index, tt.want, tt.wantIndex)
			}
		})
	}
}

func TestFindFirstReturnsSliceElement(t *testing.T) {
	negativeZero := math.Copysign(0, -1)
	got, index := FindFirst([]float64{negativeZero}, 0)
	if index != 0 || got != 0 || !math.Signbit(got) {
		t.Fatalf("FindFirst([-0], 0) = (%v, %d), want (-0, 0)", got, index)
	}
}
