package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		text, pattern string
		want          []int
	}{
		{"ababcababc", "abc", []int{2, 7}},
		{"aaaa", "aa", []int{0, 1, 2}}, // overlapping matches
		{"hello", "world", []int{}},
		{"abc", "abc", []int{0}},
		{"abc", "abcd", []int{}},
		{"abc", "", []int{}},
		{"mississippi", "issi", []int{1, 4}},
	}
	for _, c := range cases {
		if got := Search(c.text, c.pattern); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Search(%q,%q) = %v, want %v", c.text, c.pattern, got, c.want)
		}
	}
}

func TestAgainstStdlib(t *testing.T) {
	text := "the quick brown fox jumps over the lazy dog the end"
	for _, p := range []string{"the", "fox", "dog", "zzz", "q"} {
		got := Search(text, p)
		// Cross-check first occurrence with the standard library.
		want := strings.Index(text, p)
		if want == -1 && len(got) != 0 {
			t.Errorf("pattern %q: expected no matches, got %v", p, got)
		}
		if want != -1 && (len(got) == 0 || got[0] != want) {
			t.Errorf("pattern %q: first match %v, want %d", p, got, want)
		}
	}
}
