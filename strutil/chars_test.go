package strutil_test

import (
	"testing"

	"github.com/yousifsabah0/strutil/strutil"
)

func TestCharAt(t *testing.T) {
	var tests = []struct {
		str  string
		pos  int
		want string
	}{
		{"", 1, ""},
		{"h", 1, "h"},
		{"hello", -3, ""},
		{"hello", 1, "e"},
	}

	for _, test := range tests {
		got := strutil.CharAt(test.str, test.pos)
		if got != test.want {
			t.Errorf("want %q; got %q", test.want, got)
		}
	}
}

func TestCharCodeAt(t *testing.T) {
	var tests = []struct {
		str       string
		pos, want int
	}{
		{"", 1, -1},
		{"Hello", 4, 111},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			code := strutil.CharCodeAt(test.str, test.pos)
			if code != rune(test.want) {
				t.Errorf("want %q; got %q", test.want, code)
			}
		})
	}
}
