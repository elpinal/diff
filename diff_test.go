package diff

import (
	"strings"
	"testing"
)

var tests = []struct {
	a, b string
	want string
}{
	{
		"aaa aaa",
		"aaa bbb",
		Info{7, 7, 4, 'a', 'b'}.String(),
	},
	{
		"aaa",
		"aaa",
		NoDiff,
	},
	{
		"aaa",
		"aaa bbb",
		EndDiff,
	},
}

func TestDiff(t *testing.T) {
	for i, tt := range tests {
		if got := Diff(tt.a, tt.b); got != tt.want {
			t.Errorf("TestDiff(%v): got %q, want %q", i, got, tt.want)
		}
	}
}

var (
	benchBase = strings.Repeat("aaaaa bbbbb", 50)
)

func BenchmarkDiff(b *testing.B) {
	s1 := benchBase + "c"
	s2 := benchBase + "d"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(s1, s2)
	}
}
