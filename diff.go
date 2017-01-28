// Package diff provides utilities to compare strings for debugging.
package diff

import "fmt"

var (
	// NoDiff means no differences.
	NoDiff  = ""
	// EndDiff means one text contanins the other text.
	EndDiff = "diff is in the end of text"
)

// Info represents information on differences.
type Info struct {
	l1, l2 int
	i      int
	a, b   byte
}

func (i Info) String() string {
	return fmt.Sprintf("len[%v and %v]: at %v, %q and %q", i.l1, i.l2, i.i, string(i.a), string(i.b))
}

// Diff returns differences between two string.
func Diff(a, b string) string {
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return Info{len(a), len(b), i, a[i], b[i]}.String()
		}
	}
	if len(a) == len(b) {
		return NoDiff
	}
	return EndDiff
}
