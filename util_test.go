package main

import (
	"testing"
)

var distanceTests = []struct {
	in  []Coord
	out int
}{
	{[]Coord{{2, 3}, {2, 3}}, 0},
	{[]Coord{{2, 3}, {1, 3}}, 1},
	{[]Coord{{2, 3}, {2, 4}}, 1},
	{[]Coord{{2, 3}, {3, 4}}, 2},
	{[]Coord{{2, 3}, {0, 0}}, 5},
}

func TestDistance(t *testing.T) {
	for _, tt := range distanceTests {
		t.Run("distance", func(t *testing.T) {
			s := distance(tt.in[0], tt.in[1])
			if s != tt.out {
				t.Errorf("for %v got %v, want %v", tt.in, s, tt.out)
			}
		})
	}
}
