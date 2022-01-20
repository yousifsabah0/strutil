package strutil_test

import (
	"testing"

	"github.com/yousifsabah0/strutil/strutil"
)

func TestBasicSearch(t *testing.T) {
	var tests = []struct {
		str, pattern string
		want         int
	}{
		{"BBACCAADDEE", "HBB", -1},
	}

	for _, test := range tests {
		got := strutil.BasicSearch(test.str, test.pattern)
		if got != test.want {
			t.Errorf("got %q; want %q", got, test.want)
		}
	}
}
