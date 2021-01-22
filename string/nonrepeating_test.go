package main

import "testing"

func TestSubstr(t *testing.T) {
	test := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"b", 1},
		{"abcabcabcd", 4},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range test {
		actual := noRepeating(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

// go tool pprof cpu.out
// web
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual := noRepeating(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
