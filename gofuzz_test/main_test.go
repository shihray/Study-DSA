package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello", " ", "98765#"}
	for _, tc := range testcases {
		f.Add(tc) // 把 testcase 加進 seed corpus裡
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, _ := Reverse(orig)

		doubleRev, _ := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
