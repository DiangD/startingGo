package main

import "testing"

func Test_calcTriangle(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		// TODO: Add test cases.
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 57},
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d) got %d expected %d ",
				tt.a, tt.b, tt.c, actual)
		}
	}
}
