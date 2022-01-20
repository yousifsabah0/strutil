package strutil_test

import (
	"testing"

	"github.com/yousifsabah0/strutil/strutil"
)

func TestIndexOf(t *testing.T) {
	var tests = []struct {
		str, tar string
		want     int
	}{{"", "H", -1}, {"Hello", "H", 0}, {"Hello", "o", 4}, {"Hello", "", 0}}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			got := strutil.IndexOf(test.str, test.tar)
			if got != test.want {
				t.Errorf("want %q; got %q", test.want, got)
			}
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	var tests = []struct {
		str, tar string
		want     int
	}{{"", "H", -1}, {"Hello", "", 0}, {"Hello, World", "o", 8}, {"Hello, World How are u", "o", 14}}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			got := strutil.LastIndexOf(test.str, test.tar)
			if got != test.want {
				t.Errorf("want %q; got %q", test.want, got)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	var tests = []struct {
		str, char string
		want      bool
	}{
		{"Hello", "H", true},
		{"Hello", "e", false},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			got := strutil.StartsWith(test.str, test.char)
			if got != test.want {
				t.Errorf("want %v; got %v", test.want, got)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	var tests = []struct {
		str, char string
		want      bool
	}{
		{"Hello", "o", true},
		{"Hello", "e", false},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			got := strutil.EndsWith(test.str, test.char)
			if got != test.want {
				t.Errorf("want %v; got %v", test.want, got)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	var tests = []struct {
		str, char, want string
	}{
		{"Hello", "Hello", "HelloHello"},
		{"Hello", ", world", "Hello, world"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			got := strutil.Concat(test.str, test.char)
			if got != test.want {
				t.Errorf("want %v; got %v", test.want, got)
			}
		})
	}
}
