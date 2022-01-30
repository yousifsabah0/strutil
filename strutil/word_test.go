package strutil_test

import (
	"testing"

	"github.com/yousifsabah0/strutil/strutil"
)

func TestFindLongestWord(t *testing.T) {
	var tests = []struct {
		str, want string
  }{{"Helloo world", "Helloo"}, {"Helllo Woooorld", "Woooorld"}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			got := strutil.FindLongestWord(test.str, test.tar)
			if got != test.want {
				t.Errorf("want %q; got %q", test.want, got)
			}
		})
	}
}
