package main

import "testing"

func Test_lengthOfNomRepeatingSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// TODO: Add test cases.
		{"", 0},
		{"abcabcbb", 3},
	}
	for _, tt := range tests {
		if got := lengthOfNomRepeatingSubStr(tt.s); got != tt.ans {
			t.Errorf("lengthOfNomRepeatingSubStr(%v) = %v, want %v", tt.s, got, tt.ans)
		}
	}
}

func Benchmark(b *testing.B) {
	s := "黑化肥挥发会发黑黑化肥挥发会发黑黑化肥挥发会发黑"
	ans := 6
	for i := 0; i < b.N; i++ {
		res := lengthOfNomRepeatingSubStr(s)
		if ans != res {
			b.Errorf("got  %d for input %s,expected %d",
				res, s, ans)
		}
	}
}
