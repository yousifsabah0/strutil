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
		{"H", 5, "H"},
		{"Hello", 3, "l"},
		{"Hello", -3, ""},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := strutil.CharAt(test.str, test.pos)
			if str != test.want {
				t.Errorf("want %q; got %q", test.want, str)
			}
		})
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
